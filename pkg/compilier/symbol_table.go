package compilier

import (
	"fmt"
)

type SymbolScope string

const (
	LocalScope    SymbolScope = "LOCAL"
	GlobalScope   SymbolScope = "GLOBAL"
	BultinScope   SymbolScope = "BULTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FUNCTION"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
	Mut   bool
}

type SymbolTable struct {
	Outer          *SymbolTable
	store          map[string]Symbol
	numDefinitions int
	FreeSymbols    []Symbol
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	free := []Symbol{}
	return &SymbolTable{store: s, FreeSymbols: free}
}

func (s *SymbolTable) Define(name string, mut bool) (Symbol, error) {
	symbol := Symbol{
		Name:  name,
		Index: s.numDefinitions,
		Mut:   mut,
	}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	obj, ok := s.Resolve(name)
	if ok {
		if symbol.Scope == obj.Scope {
			return obj, fmt.Errorf("symbol %s is already declard", obj.Name)
		}
	}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol, nil
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BultinScope {
			return obj, ok
		}

		free := s.DefineFree(obj)
		return free, true
	}
	return obj, ok
}

func NewEncolsedTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func (s *SymbolTable) DefineBultin(index int, name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: index,
		Scope: BultinScope,
	}
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: 0,
		Scope: FunctionScope,
	}
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) DefineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)

	symbol := Symbol{
		Name:  original.Name,
		Index: len(s.FreeSymbols) - 1,
	}
	symbol.Scope = FreeScope

	s.store[original.Name] = symbol
	return symbol
}
