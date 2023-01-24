package builtins

import (
	"fmt"

	"github.com/NickDeChip/bottle-brush/pkg/object"
)

var Builtins = map[string]*object.Builtin{
	"say": say,
	"len": bblen,
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}
