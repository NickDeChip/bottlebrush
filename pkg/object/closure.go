package object

import "fmt"

type Closure struct {
	Fn   CompiledFunction
	Free []Object
}

func (c *Closure) Type() Type {
	return CLOSORE
}

func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
