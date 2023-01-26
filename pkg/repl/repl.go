package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/NickDeChip/bottlebrush/pkg/builtins"
	"github.com/NickDeChip/bottlebrush/pkg/compilier"
	"github.com/NickDeChip/bottlebrush/pkg/lexer"
	"github.com/NickDeChip/bottlebrush/pkg/object"
	"github.com/NickDeChip/bottlebrush/pkg/parser"
	"github.com/NickDeChip/bottlebrush/pkg/vm"
)

const PROMT = "=>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(scanLinesEscapable)

	variables := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compilier.NewSymbolTable()
	for i, v := range builtins.Builtins {
		symbolTable.DefineBultin(i, v.Name)
	}

	for {
		print(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		prog := p.ParseProgram()
		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
			continue
		}

		comp := compilier.NewWithState(symbolTable, variables)
		err := comp.Compile(prog)
		if err != nil {
			fmt.Fprintf(out, "Darn! Compilation failed!\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		variables = code.Variables

		machine := vm.NewWithGlobleStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Darn! executing bytecode failed!\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func scanLinesEscapable(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.LastIndexByte(data, '\n'); i >= 0 {
		if wasNewLibeEscaped(data[0:i]) {
			// New line was escaped, get more data
			return 0, nil, nil
		}
		return i + 1, dropCRBS(data[0:i]), nil
	}

	if atEOF {
		return len(data), dropCRBS(data), nil
	}

	return 0, nil, nil
}

func wasNewLibeEscaped(data []byte) bool {
	if len(data) > 0 {
		if data[len(data)-1] == '\r' {
			return wasNewLibeEscaped(dropCR(data))
		}
		if data[len(data)-1] == '\\' {
			return true
		}
	}
	return false
}

func dropCRBS(data []byte) []byte {
	if len(data) > 0 {
		return bytes.ReplaceAll(
			bytes.ReplaceAll(data, []byte("\r"), []byte("")),
			[]byte("\\\n"),
			[]byte("\n"),
		)
	}
	return data
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
