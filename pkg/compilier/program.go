package compilier

import "github.com/NickDeChip/bottlebrush/pkg/ast"

func compileProgram(node *ast.Program, c *Compilier) error {
	for _, s := range node.Statements {
		err := c.Compile(s)
		if err != nil {
			return err
		}
	}

	return nil
}
