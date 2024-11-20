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
	l.Pos++
	if l.Pos >= len(l.Source) {
		return fmt.Errorf("EOF")
	}
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
func (l *Lexer) GetToken() {}
