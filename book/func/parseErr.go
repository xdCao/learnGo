package main

import "fmt"

func main() {

	var e error

	e = newParseError("main.go", 1)

	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("filename: %s , line: %d\n", detail.filename, detail.line)
	}

}

type ParseError struct {
	filename string
	line     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s : %d", e.filename, e.line)
}

func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
