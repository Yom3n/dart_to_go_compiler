package main

import "fmt"

func main() {
	source := "var a = 123;"
	if len(source) == 0 {
		return
	}
	lexer := MakeLexer(source)
	token, err := lexer.GetToken()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	for token.Kind != EOF {
		fmt.Println(token.Kind)
		token, err = lexer.GetToken()
		if err != nil {
			fmt.Printf("Error getting token: %s", err.Error())
			return
		}
	}
}
