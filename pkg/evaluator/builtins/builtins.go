package builtins

import "github.com/NickDeChip/bottle-brush/pkg/object"

var Builtins = map[string]*object.Builtin{
	"say": say,
}
