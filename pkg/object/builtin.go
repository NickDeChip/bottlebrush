package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type {
	return BUILTIN
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}
