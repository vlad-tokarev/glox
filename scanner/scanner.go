package scanner

import (
	"errors"
	"fmt"
	"github.com/vlad-tokarev/glox/error_reporter"
	"strconv"
	"unicode"
)

var (
	ErrDone                = errors.New("scanner: done")
	errUnexpectedCharacter = errors.New("scanner: unexpected character")
	errIgnoredCharacter    = errors.New("scanner: ignored character")
)

type Literal struct {
	Number float64
	String string
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal Literal
	Line    int64
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
}

type Scanner struct {
	source   []rune
	hasError bool

	start   int64
	current int64
	line    int64
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: []rune(source), line: 1}
}

// Next returns next token.
// Returns ErrDone if no more tokens.
func (s *Scanner) Next() (Token, error) {

	token, err := s.scanToken()

	for {
		switch err {
		case errIgnoredCharacter, errUnexpectedCharacter:
			token, err = s.scanToken()
		default:
			return token, err
		}

	}
}

func (s *Scanner) scanToken() (Token, error) {

	if s.isAtEnd() {
		return Token{Type: EOF}, ErrDone
	}
	s.start = s.current

	c := s.advance()
	switch c {
	case '(':
		return Token{Type: LeftParen, Lexeme: string(c), Line: s.line}, nil
	case ')':
		return Token{Type: RightParen, Lexeme: string(c), Line: s.line}, nil
	case '{':
		return Token{Type: LeftBrace, Lexeme: string(c), Line: s.line}, nil
	case '}':
		return Token{Type: RightBrace, Lexeme: string(c), Line: s.line}, nil
	case ',':
		return Token{Type: Comma, Lexeme: string(c), Line: s.line}, nil
	case '.':
		return Token{Type: Dot, Lexeme: string(c), Line: s.line}, nil
	case '-':
		return Token{Type: Minus, Lexeme: string(c), Line: s.line}, nil
	case '+':
		return Token{Type: Plus, Lexeme: string(c), Line: s.line}, nil
	case ';':
		return Token{Type: Semicolon, Lexeme: string(c), Line: s.line}, nil
	case '*':
		return Token{Type: Star, Lexeme: string(c), Line: s.line}, nil
	case '!':
		if s.match('=') {
			return Token{Type: BangEqual, Lexeme: "!=", Line: s.line}, nil
		} else {
			return Token{Type: Bang, Lexeme: string(c), Line: s.line}, nil
		}
	case '=':
		if s.match('=') {
			return Token{Type: EqualEqual, Lexeme: "==", Line: s.line}, nil
		} else {
			return Token{Type: Equal, Lexeme: string(c), Line: s.line}, nil
		}
	case '<':
		if s.match('=') {
			return Token{Type: LessEqual, Lexeme: "<=", Line: s.line}, nil
		} else {
			return Token{Type: Less, Lexeme: string(c), Line: s.line}, nil
		}
	case '>':
		if s.match('=') {
			return Token{Type: GreaterEqual, Lexeme: ">=", Line: s.line}, nil
		} else {
			return Token{Type: Greater, Lexeme: string(c), Line: s.line}, nil
		}
	case '/':
		if s.match('/') {
			for {
				if s.isAtEnd() {
					break
				}
				if s.peek() == '\n' {
					break
				}
				s.advance()
			}
		} else {
			return Token{Type: Slash, Lexeme: string(c), Line: s.line}, nil
		}
	case ' ', '\r', '\t':
		return Token{}, errIgnoredCharacter
	case '\n':
		s.line++
		return Token{}, errIgnoredCharacter
	case '"':
		return s.scanString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return s.scanNumber()
	default:
		if unicode.IsLetter(c) || c == '_' {
			return s.scanIdentifier()
		}
	}

	error_reporter.Print(int(s.line), fmt.Sprintf("Unexpected character: %s", string(c)))
	return Token{}, errUnexpectedCharacter
}

func (s *Scanner) scanString() (Token, error) {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		error_reporter.Print(int(s.line), "Unterminated string.")
		return Token{}, errUnexpectedCharacter
	}

	// The closing ".
	s.advance()

	value := string(s.source[s.start+1 : s.current-1])

	return Token{Type: String, Lexeme: value, Line: s.line, Literal: Literal{
		String: value,
	}}, nil

}

func (s *Scanner) scanNumber() (Token, error) {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
		// consume the "."
		s.advance()

		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	value := string(s.source[s.start:s.current])
	token := Token{
		Type:    Number,
		Lexeme:  value,
		Literal: Literal{},
		Line:    s.line,
	}
	var err error
	token.Literal.Number, err = strconv.ParseFloat(value, 64)
	return token, err
}

func (s *Scanner) scanIdentifier() (Token, error) {
	for unicode.IsLetter(s.peek()) || unicode.IsDigit(s.peek()) || s.peek() == '_' {
		s.advance()
	}

	value := string(s.source[s.start:s.current])

	token := Token{Type: Identifier, Lexeme: string(s.source[s.start:s.current]), Line: s.line}

	switch value {
	case "and":
		token.Type = And
	case "class":
		token.Type = Class
	case "else":
		token.Type = Else
	case "false":
		token.Type = False
	case "for":
		token.Type = For
	case "fun":
		token.Type = Fun
	case "if":
		token.Type = If
	case "nil":
		token.Type = Nil
	case "or":
		token.Type = Or
	case "print":
		token.Type = Print
	case "return":
		token.Type = Return
	case "super":
		token.Type = Super
	case "this":
		token.Type = This
	case "true":
		token.Type = True
	case "var":
		token.Type = Var
	case "while":
		token.Type = While
	}

	return token, nil
}
