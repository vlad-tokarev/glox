package parser

import (
	"github.com/vlad-tokarev/glox/scanner"
	"log"
)

func (p *Parser) match(tts ...scanner.TokenType) bool {
	for _, tt := range tts {
		if p.check(tt) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() {
	if !p.isAtEnd() {
		token, err := p.scanner.Next()
		if err != nil {
			log.Fatalf("Could not scan: %s", err)
		}
		p.prev = p.current
		p.current = token
	}
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == scanner.EOF
}

func (p *Parser) peek() scanner.Token {
	return p.current
}

func (p *Parser) previous() scanner.Token {
	return p.prev
}
