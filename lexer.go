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
func (l *Lexer) Abort(message string) {
	msg := "Lexing error. " + message
	panic(msg)
}

// Skip whitespace except newlines, which we will use to indicate the end of a statement
func (l *Lexer) SkipWhiteSpace() {
	char := l.CurrentChar()
	for char == " " || char == "\t" || char == "\r" {
		fmt.Println("Skipping char: ")
		fmt.Println(char)
		err := l.NextChar()
		if err != nil {
			break
		}
	}
}

// Skip comments in code
func (l *Lexer) SkipComment() {}

// Return next token
func (l *Lexer) GetToken() (Token, error) {
	l.SkipWhiteSpace()
	var token Token

	next, err := l.Peek()
	switch char := l.CurrentChar(); char {
	case "+":
		token = Token{Text: char, Kind: PLUS}
	case "-":
		token = Token{Text: char, Kind: MINUS}
	case "=":
		if err != nil && next == "=" {
			lastChar := l.CurrentChar()
			l.NextChar()
			token = Token{Text: lastChar + next, Kind: EQEQ}
		} else {
			token = Token{Text: char, Kind: EQ}
		}
	case "*":
		token = Token{Text: char, Kind: ASTERISK}
	case "/":
		token = Token{Text: char, Kind: SLASH}
	case "\n":
		token = Token{Text: char, Kind: NEW_LINE}
	case ">":
		if err == nil && next == "=" {
			lastChar := l.CurrentChar()
			l.NextChar()
			token = Token{Text: lastChar + next, Kind: GTEQ}
		} else {
			token = Token{Text: char, Kind: GT}
		}
	case "<":
		if err == nil && next == "=" {
			lastChar := l.CurrentChar()
			l.NextChar()
			token = Token{Text: lastChar + next, Kind: LTEQ}
		} else {
			token = Token{Text: char, Kind: LT}
		}
	case "!":
		if err == nil && next == "=" {
			lastChar := l.CurrentChar()
			l.NextChar()
			token = Token{Text: lastChar + next, Kind: NOTEQ}
		} else {
			l.Abort("Expected !=, got !")
		}
	default:
		l.Abort("Unknown token: '" + char + "'")
	}

	err = l.NextChar()
	if err != nil {
		return Token{Kind: EOF}, nil
	}
	return token, nil
}
