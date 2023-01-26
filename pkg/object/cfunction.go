package object

import (
	"fmt"

	"github.com/NickDeChip/bottlebrush/pkg/opcode"
)

type CompiledFunction struct {
	Instructions opcode.Instructions
	NumLocals    int
	NumParameter int
}

func (cf *CompiledFunction) Type() Type {
	return CFUNCTION
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompileFunction[%p]", cf)
}
