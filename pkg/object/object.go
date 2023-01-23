package object

type Type string

const (
	STRING  Type = "STRING"
	INTEGER Type = "INTEGER"
	FLOAT   Type = "FLOAT"
	BOOL    Type = "BOOL"
	ERROR   Type = "ERROR"
	FUNTION Type = "FUNTION"
	RETURN  Type = "RETURN"
	BUILTIN Type = "BUILTIN"
	NULL    Type = "NULL"
)

type Object interface {
	Type() Type
	Inspect() string
}
