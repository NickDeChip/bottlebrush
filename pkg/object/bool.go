package object

import "fmt"

type Bool struct {
	Value bool
}

func (b *Bool) Type() Type {
	return BOOL
}

func (b *Bool) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
