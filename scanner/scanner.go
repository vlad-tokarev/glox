package scanner

import (
	"errors"
	"fmt"
	"github.com/vlad-tokarev/glox/error_reporter"
)

var (
	ErrDone                = errors.New("scanner: done")
	errUnexpectedCharacter = errors.New("scanner: unexpected character")
	errIgnoredCharacter    = errors.New("scanner: ignored character")
)

type LiteralType = string

const (
	LiteralNumber = "LITERAL_NUMBER"
	LiteralString = "LITERAL_STRING"
	LiteralNil    = "LITERAL_NIL"
)

type Literal struct {
	Type   LiteralType
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
	}

	error_reporter.Print(int(s.line), "Unexpected character.")
	return Token{}, errUnexpectedCharacter
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= int64(len(s.source))
}

func (s *Scanner) advance() rune {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
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

	return Token{Type: LiteralString, Lexeme: value, Line: s.line, Literal: Literal{
		Type:   LiteralString,
		String: value,
	}}, nil

}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}
