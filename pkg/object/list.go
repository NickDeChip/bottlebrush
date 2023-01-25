package object

import (
	"bytes"
	"strings"
)

type List struct {
	Elements []Object
}

func (l *List) Type() Type {
	return LIST
}

func (l *List) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range l.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
