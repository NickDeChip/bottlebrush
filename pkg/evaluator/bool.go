package evaluator

import "github.com/NickDeChip/bottle-brush/pkg/object"

func evalBool(input bool) *object.Bool {
	if input {
		return TRUE
	}
	return FALSE
}
