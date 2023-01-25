package object

import "fmt"

type Float struct {
	Value float64
}

func (f *Float) Type() Type {
	return FLOAT
}

func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}
