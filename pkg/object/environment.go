package object

type Var struct {
	Object Object
	Mut    bool
}

type Environment struct {
	store map[string]Var
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Var)
	return &Environment{
		store: s,
		outer: nil,
	}
}

func NewEncolsedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Var, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) GetVar(name string) (Var, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Var) Var {
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Set(name, val)
	}
	e.store[name] = val
	return val
}

func (e *Environment) SetVar(name string, val Var) Var {
	e.store[name] = val
	return val
}
