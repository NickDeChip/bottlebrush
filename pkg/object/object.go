package object

type Type string

const (
	STRING Type = "STRING"
)

type Object interface {
	Type() Type
	Inspect() string
}
