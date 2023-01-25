package repl

import (
	"bufio"
	"bytes"
	"io"

	"github.com/NickDeChip/bottlebrush/pkg/evaluator"
	"github.com/NickDeChip/bottlebrush/pkg/lexer"
	"github.com/NickDeChip/bottlebrush/pkg/object"
	"github.com/NickDeChip/bottlebrush/pkg/parser"
)

const PROMT = "=>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(scanLinesEscapable)

	env := object.NewEnvironment()

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

		res := evaluator.Eval(prog, env)
		if res != nil {
			io.WriteString(out, res.Inspect())
			io.WriteString(out, "\n")
		}
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
