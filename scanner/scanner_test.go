package scanner

import (
	"testing"
)

// table test
func TestScanner(t *testing.T) {
	type test struct {
		name  string
		input string
		want  []Token
	}

	tests := []test{
		{name: "grouping stuff", input: "(( )){}", want: []Token{
			{Type: LeftParen, Lexeme: "(", Line: 1},
			{Type: LeftParen, Lexeme: "(", Line: 1},
			{Type: RightParen, Lexeme: ")", Line: 1},
			{Type: RightParen, Lexeme: ")", Line: 1},
			{Type: LeftBrace, Lexeme: "{", Line: 1},
			{Type: RightBrace, Lexeme: "}", Line: 1},
		}},
		{name: "operators", input: "!*+-/=<> <= ==", want: []Token{
			{Type: Bang, Lexeme: "!", Line: 1},
			{Type: Star, Lexeme: "*", Line: 1},
			{Type: Plus, Lexeme: "+", Line: 1},
			{Type: Minus, Lexeme: "-", Line: 1},
			{Type: Slash, Lexeme: "/", Line: 1},
			{Type: Equal, Lexeme: "=", Line: 1},
			{Type: Less, Lexeme: "<", Line: 1},
			{Type: Greater, Lexeme: ">", Line: 1},
			{Type: LessEqual, Lexeme: "<=", Line: 1},
			{Type: EqualEqual, Lexeme: "==", Line: 1},
		}},
		{name: "string_and_numbers", input: "231231\"Привет\"312.21", want: []Token{
			{Type: Number, Lexeme: "231231", Line: 1, Literal: Literal{Number: 231231}},
			{Type: String, Lexeme: "Привет", Line: 1, Literal: Literal{String: "Привет"}},
			{Type: Number, Lexeme: "312.21", Line: 1, Literal: Literal{Number: 312.21}},
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewScanner(tc.input)
			for _, want := range tc.want {
				got, err := s.Next()
				if err != nil {
					t.Errorf("Exepcted token: %+v, got error %+v", want, err)
					continue
				}
				if got != want {
					t.Errorf("Expected token %+v, got %+v", want, got)
				}
			}
		})
	}
}
