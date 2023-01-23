package builtins

import "github.com/NickDeChip/bottle-brush/pkg/object"

var say = &object.Builtin{
	Fn: func(args ...object.Object) object.Object {
		for _, arg := range args {
			print(arg.Inspect())
		}
		println()
		return nil
	},
}
