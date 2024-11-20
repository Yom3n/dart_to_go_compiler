package main

import "fmt"

func main() {
	source := "var a = 123;"
	if len(source) == 0 {
		return
	}
	lexer := MakeLexer(source)
	var err error = nil
	for ; err == nil; _, err = lexer.Peek() {
		fmt.Println(lexer.CurrentChar())
		lexer.NextChar()
	}
}
