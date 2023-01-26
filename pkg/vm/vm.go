package vm

import (
	"fmt"

	"github.com/NickDeChip/bottlebrush/pkg/builtins"
	"github.com/NickDeChip/bottlebrush/pkg/compilier"
	"github.com/NickDeChip/bottlebrush/pkg/object"
	"github.com/NickDeChip/bottlebrush/pkg/opcode"
)

const (
	StackSize   = 2048
	GlobalsSize = 65536
	MaxFrames   = 1024
)

var (
	True  = &object.Bool{Value: true}
	False = &object.Bool{Value: false}
	Null  = &object.Null{}
)

type VM struct {
	variables []object.Object

	stack []object.Object
	sp    int

	globals []object.Object

	frames     []*Frame
	frameIndex int
}

func New(byteCode *compilier.ByteCode) *VM {
	mainFn := &object.CompiledFunction{
		Instructions: byteCode.Instrctinos,
	}
	mainClosure := object.Closure{
		Fn: *mainFn,
	}
	mainFrame := NewFrame(&mainClosure, 0)

	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame
	return &VM{
		variables: byteCode.Variables,

		stack: make([]object.Object, StackSize),
		sp:    0,

		globals: make([]object.Object, GlobalsSize),

		frames:     frames,
		frameIndex: 1,
	}
}

func NewWithGlobleStore(bytecode *compilier.ByteCode, s []object.Object) *VM {
	vm := New(bytecode)
	vm.globals = s
	return vm
}

func (vm *VM) StackTop() object.Object {
	if vm.sp != 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) push(o object.Object) error {
	if vm.sp > StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) Run() error {
	var ip int
	var ins opcode.Instructions
	var op opcode.Opcode

	for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
		vm.currentFrame().ip++

		ip = vm.currentFrame().ip
		ins = vm.currentFrame().Instructions()
		op = opcode.Opcode(ins[ip])
		switch op {
		case opcode.OpVariable:
			varIndex := opcode.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			err := vm.push(vm.variables[varIndex])
			if err != nil {
				return err
			}

		case opcode.OpPop:
			vm.pop()

		case opcode.OpSetGlobal:
			globleIndex := opcode.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			vm.globals[globleIndex] = vm.pop()

		case opcode.OpGetGlobal:
			globleIndex := opcode.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2
			err := vm.push(vm.globals[globleIndex])
			if err != nil {
				return err
			}

		case opcode.OpCall:
			numArgs := opcode.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip++

			err := vm.exucuteCall(int(numArgs))
			if err != nil {
				return err
			}

		case opcode.OpReturnValue:
			returnValue := vm.pop()

			frame := vm.popFrame()
			vm.sp = frame.basePointer - 1

			err := vm.push(returnValue)
			if err != nil {
				return err
			}

		case opcode.OpGetBuiltin:
			builtinIndex := opcode.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip++

			definitition := builtins.Builtins[builtinIndex]

			err := vm.push(definitition.Builtin)
			if err != nil {
				return err
			}

		case opcode.OpGetFree:
			freeIndex := opcode.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip++

			currentClosure := vm.currentFrame().cl
			err := vm.push(currentClosure.Free[freeIndex])
			if err != nil {
				return err
			}

		case opcode.OpCurrentClosure:
			currentClosure := vm.currentFrame().cl
			err := vm.push(currentClosure)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.frameIndex-1]
}

func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.frameIndex] = f
	vm.frameIndex++
}

func (vm *VM) popFrame() *Frame {
	vm.frameIndex--
	return vm.frames[vm.frameIndex]
}

func (vm *VM) exucuteCall(numArgs int) error {
	callee := vm.stack[vm.sp-1-numArgs]
	switch callee := callee.(type) {
	case *object.Closure:
		// TODO: Surrport
		return fmt.Errorf("closure not surrported")

	case *object.Builtin:
		return vm.CallBuiltin(callee, numArgs)

	default:
		return fmt.Errorf("calling non closure and non-builtin")
	}
}

func (vm *VM) CallBuiltin(builtin *object.Builtin, numArgs int) error {
	args := vm.stack[vm.sp-numArgs : vm.sp]

	result := builtin.Fn(args...)
	vm.sp = vm.sp - numArgs - 1
	if result != nil {
		vm.push(result)
	} else {
		vm.push(Null)
	}

	return nil
}
