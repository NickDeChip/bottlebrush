package builtins

import (
	"fmt"

	"github.com/NickDeChip/bottlebrush/pkg/object"
)

type Builtin struct {
	Name    string
	Builtin *object.Builtin
}

var Builtins = append(
	[]Builtin{},
	stdBuiltin...,
)

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}
