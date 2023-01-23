package object

import (
	"bytes"
)

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type {
	return RETURN
}

func (rv *ReturnValue) Inspect() string {
	var out bytes.Buffer

	out.WriteString("return ")
	out.WriteString(rv.Value.Inspect())

	return out.String()
}
