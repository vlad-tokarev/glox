package scanner

type TokenType = string

// Single-character tokens.
const (
	LeftParen  TokenType = "LEFT_PAREN"
	RightParen TokenType = "RIGHT_PAREN"
	LeftBrace  TokenType = "LEFT_BRACE"
	RightBrace TokenType = "RIGHT_BRACE"
	Comma      TokenType = "COMMA"
	Dot        TokenType = "DOT"
	Minus      TokenType = "MINUS"
	Plus       TokenType = "PLUS"
	Semicolon  TokenType = "SEMICOLON"
	Slash      TokenType = "SLASH"
	Star       TokenType = "STAR"
)

// One or two character tokens.
const (
	Bang         TokenType = "BANG"
	BangEqual    TokenType = "BANG_EQUAL"
	Equal        TokenType = "EQUAL"
	EqualEqual   TokenType = "EQUAL_EQUAL"
	Greater      TokenType = "GREATER"
	GreaterEqual TokenType = "GREATER_EQUAL"
	Less         TokenType = "LESS"
	LessEqual    TokenType = "LESS_EQUAL"
)

// Literals.
const (
	Identifier TokenType = "IDENTIFIER"
	String     TokenType = "STRING"
	Number     TokenType = "NUMBER"
)

// Keywords
const (
	And    TokenType = "AND"
	Class  TokenType = "CLASS"
	Else   TokenType = "ELSE"
	False  TokenType = "FALSE"
	Fun    TokenType = "FUN"
	For    TokenType = "FOR"
	If     TokenType = "IF"
	Nil    TokenType = "NIL"
	Or     TokenType = "OR"
	Print  TokenType = "PRINT"
	Return TokenType = "RETURN"
	Super  TokenType = "SUPER"
	This   TokenType = "THIS"
	True   TokenType = "TRUE"
	Var    TokenType = "VAR"
	While  TokenType = "WHILE"
)

const EOF TokenType = "EOF"
