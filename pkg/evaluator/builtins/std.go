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

var bblen = &object.Builtin{
	Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("expected one argument of type %s, got %d args", object.STRING, len(args))
		}
		if val, ok := args[0].(*object.String); ok {
			return &object.Integer{
				Value: int32(len(val.Value)),
			}
		}
		return newError("expected %s, got %s", object.STRING, args[0].Type())
	},
}
