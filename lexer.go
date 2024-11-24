package main

import (
	"fmt"
)

type Lexer struct {
	Source string
	Pos    int
}

func MakeLexer(source string) Lexer {
	l := Lexer{
		Source: source + "\n",
		Pos:    -1,
	}
	l.NextChar()
	return l
}

func (l *Lexer) CurrentChar() string {
	return string(l.Source[l.Pos])
}

// Process the next character
func (l *Lexer) NextChar() error {
	// fmt.Println("Source len:")
	// fmt.Println(len(l.Source))
	// fmt.Println("Current pos:")
	// fmt.Println(l.Pos)
	// fmt.Println("------------------------")
	if l.Pos+1 >= len(l.Source) {
		return fmt.Errorf("EOF")
	}
	l.Pos++
	return nil
}

// Returns forhead character without moving to it
func (l *Lexer) Peek() (string, error) {
	if l.Pos+1 >= len(l.Source) {
		return "", fmt.Errorf("EOF")
	}
	return string(l.Source[l.Pos+1]), nil
}

// Invalid token found
func (l *Lexer) Abort() {}

// Skip whitespace except newlines, which we will use to indicate the end of a statement
func (l *Lexer) SkipWhiteSpace() {}

// Skip comments in code
func (l *Lexer) SkipComment() {}

// Return next token
func (l *Lexer) GetToken() (Token, error) {
	var token Token
	switch char := l.CurrentChar(); char {
	case "+":
		token = Token{Text: char, Kind: PLUS}
	case "-":
		token = Token{Text: char, Kind: MINUS}
	case "=":
		token = Token{Text: char, Kind: EQ}
	case "*":
		token = Token{Text: char, Kind: ASTERISK}
	case "/":
		token = Token{Text: char, Kind: SLASH}
	default:
		// return Token{}, fmt.Errorf("Unknown token: %s", char)
	}

	err := l.NextChar()
	if err != nil {
		return Token{Kind: EOF}, nil
	}
	return token, nil
}
