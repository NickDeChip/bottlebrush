package main

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/NickDeChip/bottlebrush/pkg/evaluator"
	"github.com/NickDeChip/bottlebrush/pkg/lexer"
	"github.com/NickDeChip/bottlebrush/pkg/object"
	"github.com/NickDeChip/bottlebrush/pkg/parser"
	"github.com/NickDeChip/bottlebrush/pkg/repl"
)

var reset = "\033[0m"
var red = "\033[31m"

func init() {
	if runtime.GOOS == "windows" {
		red = ""
		reset = ""
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		println(red + "Welcome to the Bottlebrush Exsperience!" + reset)
		repl.Start(os.Stdin, os.Stdout)
	}

	data, err := ioutil.ReadFile(args[0])

	if err != nil {
		println("could not read file!")
	}

	l := lexer.New(string(data))
	p := parser.New(l)
	prog := p.ParseProgram()

	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return
	}

	env := object.NewEnvironment()
	res := evaluator.Eval(prog, env)

	runtimeError, ok := res.(*object.Error)
	if ok {
		println(runtimeError.Inspect())
	}

}
