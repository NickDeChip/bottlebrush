package evaluator

import (
	"fmt"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}

	args := evalExpresions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyFunction(function, args, node.Line(), node.Col())
}

func applyFunction(fn object.Object, args []object.Object, line int, col int) object.Object {
	switch fn := fn.(type) {
	case *object.Fn:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		if res := fn.Fn(args...); res != nil {
			if err, ok := res.(*object.Error); ok {
				err.Message += fmt.Sprintf("; line=%d; col=%d", line, col)
			}
			return res
		}
		return NULL
	default:
		return newError("not a function: %s; line=%d; col=%d", fn.Type(), line, col)
	}

}

func extendFunctionEnv(fn *object.Fn, args []object.Object) *object.Environment {
	env := object.NewEncolsedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.SetVar(param.Value, object.Var{
			Object: args[paramIdx],
			Mut:    true,
		})
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	if err, ok := obj.(*object.Error); ok {
		return err
	}
	return NULL
}
