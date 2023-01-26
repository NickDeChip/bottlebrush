package compilier

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/opcode"
)

func compileVar(node *ast.VarStatement, c *Compilier) error {
	symbol, err := c.symbolTable.Define(node.Name.Value, node.Mut)
	if err != nil {
		return err
	}
	err = c.Compile(node.Value)
	if err != nil {
		return err
	}

	if symbol.Scope == GlobalScope {
		c.emit(opcode.OpSetGlobal, symbol.Index)
	} else {
		c.emit(opcode.OpSetLocal, symbol.Index)
	}

	return nil
}
