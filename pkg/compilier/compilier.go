package compilier

import (
	"fmt"

	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/builtins"
	"github.com/NickDeChip/bottlebrush/pkg/object"
	"github.com/NickDeChip/bottlebrush/pkg/opcode"
)

type ExmittedInstruction struct {
	Opcode   opcode.Opcode
	Position int
}

type CompilationScope struct {
	instructions        opcode.Instructions
	lastInstruction     ExmittedInstruction
	previousInstruction ExmittedInstruction
}

type Compilier struct {
	variables []object.Object

	symbolTable *SymbolTable

	scopes     []CompilationScope
	scopeIndex int
}

func New() *Compilier {
	mainScope := CompilationScope{
		instructions:        opcode.Instructions{},
		lastInstruction:     ExmittedInstruction{},
		previousInstruction: ExmittedInstruction{},
	}

	symbolTable := NewSymbolTable()

	for i, v := range builtins.Builtins {
		symbolTable.DefineBultin(i, v.Name)
	}

	return &Compilier{
		variables:   []object.Object{},
		symbolTable: symbolTable,
		scopes:      []CompilationScope{mainScope},
		scopeIndex:  0,
	}
}

func (c *Compilier) Compile(node ast.Node) error {
	switch node := node.(type) {

	case *ast.Program:
		return compileProgram(node, c)

	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(opcode.OpPop)

	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined variables %s; line=%d; col%d", node.Value, node.Token.Line, node.Token.Col)
		}
		c.loadSymbol(symbol)

	case *ast.VarStatement:
		return compileVar(node, c)

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(opcode.OpVariable, c.addVariable(integer))

	case *ast.CallExpression:
		err := c.Compile(node.Function)
		if err != nil {
			return err
		}
		for _, a := range node.Arguments {
			err := c.Compile(a)
			return err
		}
		c.emit(opcode.OpCall, len(node.Arguments))

	}
	return nil
}

type ByteCode struct {
	Instrctinos opcode.Instructions
	Variables   []object.Object
}

func (c *Compilier) Bytecode() *ByteCode {
	return &ByteCode{
		Instrctinos: c.currrentInstrctions(),
		Variables:   c.variables,
	}
}

func (c *Compilier) addVariable(obj object.Object) int {
	c.variables = append(c.variables, obj)
	return len(c.variables) - 1
}

func (c *Compilier) emit(op opcode.Opcode, operands ...int) int {
	ins := opcode.Make(op, operands...)
	pos := c.addInstructions(ins)

	c.setLastInstruction(op, pos)

	return pos
}

func (c *Compilier) addInstructions(ins []byte) int {
	posNewInstruction := len(c.currrentInstrctions())
	updateInstruction := append(c.currrentInstrctions(), ins...)

	c.scopes[c.scopeIndex].instructions = updateInstruction

	return posNewInstruction
}

func (c *Compilier) setLastInstruction(op opcode.Opcode, pos int) {
	previous := c.scopes[c.scopeIndex].lastInstruction
	last := ExmittedInstruction{Opcode: op, Position: pos}

	c.scopes[c.scopeIndex].previousInstruction = previous
	c.scopes[c.scopeIndex].lastInstruction = last
}

func (c *Compilier) lastInstruction(op opcode.Opcode) bool {
	if len(c.currrentInstrctions()) == 0 {
		return false
	}
	return c.scopes[c.scopeIndex].lastInstruction.Opcode == op
}

func (c *Compilier) RemoveLastPop() {
	last := c.scopes[c.scopeIndex].lastInstruction
	previous := c.scopes[c.scopeIndex].lastInstruction

	old := c.currrentInstrctions()
	new := old[:last.Position]

	c.scopes[c.scopeIndex].instructions = new
	c.scopes[c.scopeIndex].lastInstruction = previous
}

func (c *Compilier) replaceInstruction(pos int, newInstruction []byte) {
	ins := c.currrentInstrctions()

	for i := 0; i < len(newInstruction); i++ {
		ins[pos+1] = newInstruction[i]
	}
}

func (c *Compilier) changeOpernd(opPos int, operand int) {
	op := opcode.Opcode(c.currrentInstrctions()[opPos])
	NewInstruction := opcode.Make(op, operand)

	c.replaceInstruction(opPos, NewInstruction)
}

func NewWithState(s *SymbolTable, variables []object.Object) *Compilier {
	compilier := New()
	compilier.symbolTable = s
	compilier.variables = variables
	return compilier
}

func (c *Compilier) currrentInstrctions() opcode.Instructions {
	return c.scopes[c.scopeIndex].instructions
}

func (c *Compilier) enterScope() {
	scope := CompilationScope{
		instructions:        opcode.Instructions{},
		lastInstruction:     ExmittedInstruction{},
		previousInstruction: ExmittedInstruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.scopeIndex++

	c.symbolTable = NewEncolsedTable(c.symbolTable)
}

func (c *Compilier) leaveScope() opcode.Instructions {
	instructions := c.currrentInstrctions()

	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	c.symbolTable = c.symbolTable.Outer

	return instructions
}

func (c *Compilier) replaceLastPopWthReturn() {
	lastPos := c.scopes[c.scopeIndex].lastInstruction.Position

	c.replaceInstruction(lastPos, opcode.Make(opcode.OpReturnValue))

	c.scopes[c.scopeIndex].lastInstruction.Opcode = opcode.OpReturnValue
}

func (c *Compilier) loadSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.emit(opcode.OpGetGlobal, s.Index)
	case LocalScope:
		c.emit(opcode.OpGetLocal, s.Index)
	case BultinScope:
		c.emit(opcode.OpGetBuiltin, s.Index)
	case FreeScope:
		c.emit(opcode.OpGetFree, s.Index)
	case FunctionScope:
		c.emit(opcode.OpCurrentClosure, s.Index)
	}
}
