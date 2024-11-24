package main

type TokenKind int64

const (
	EOF      = -1
	UNKNOWN  = 0
	NEW_LINE = 1
	NUMBER   = 2
	IDENT    = 3
	STRING   = 4
	// Keywords
	VOID   = 101
	PRINT  = 102
	VAR    = 103
	IF     = 104
	RETURN = 105
	WHILE  = 106
	// Operators
	EQ       = 201
	PLUS     = 202
	MINUS    = 203
	ASTERISK = 204
	SLASH    = 205
	EQEQ     = 206
	NOTEQ    = 207
	LT       = 208
	LTEQ     = 209
	GT       = 210
	GTEQ     = 211
)

func (tk TokenKind) String() string {
	switch tk {
	case EOF:
		return "EOF"
	case UNKNOWN:
		return "Unknown"
	case NEW_LINE:
		return "NewLine"
	case NUMBER:
		return "Number"
	case IDENT:
		return "Ident"
	case STRING:
		return "String"
	case VOID:
		return "Void"
	case PRINT:
		return "Print"
	case VAR:
		return "Var"
	case IF:
		return "If"
	case RETURN:
		return "Return"
	case WHILE:
		return "While"
	case EQ:
		return "Equal"
	case PLUS:
		return "Plus"
	case MINUS:
		return "Minus"
	case ASTERISK:
		return "Asterisk"
	case SLASH:
		return "Slash"
	case EQEQ:
		return "EqualEqual"
	case NOTEQ:
		return "NotEqual"
	case LT:
		return "LessThan"
	case LTEQ:
		return "LessThanOrEqual"
	case GT:
		return "GreaterThan"
	case GTEQ:
		return "GreaterThanOrEqual"
	default:
		return "Unknown"
	}
}

type Token struct {
	Text string
	Kind TokenKind
}
