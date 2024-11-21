package main

type TokenKind int64

const (
	UNKNOWN = -1
	LEW_LINE = 0
	NUMBER = 1
	IDENT = 2
	STRING = 3
	// Keywords
	VOID = 101
	PRINT = 102
	VAR = 103
	IF = 104
	RETURN = 105
	WHILE = 106
	// Operators
	EQ = 201  
	PLUS = 202
	MINUS = 203
	ASTERISK = 204
	SLASH = 205
	EQEQ = 206
	NOTEQ = 207
	LT = 208
	LTEQ = 209
	GT = 210
	GTEQ = 211

)

type Token struct{
	Text string
	Kind TokenKind
}