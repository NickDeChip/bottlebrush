package evaluator

import "github.com/NickDeChip/bottlebrush/pkg/object"

func evalBool(input bool) *object.Bool {
	if input {
		return TRUE
	}
	return FALSE
}
