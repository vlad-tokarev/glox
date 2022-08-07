package parser

import (
	"github.com/vlad-tokarev/glox/expr"
	"github.com/vlad-tokarev/glox/scanner"
	"log"
)

type Parser struct {
	scanner *scanner.Scanner

	current scanner.Token
	prev    scanner.Token
}

func NewParser(scanner *scanner.Scanner, current scanner.Token, prev scanner.Token) *Parser {
	current, err := scanner.Next()
	if err != nil {
		log.Fatalf("Could not scan: %s", err)
	}
	return &Parser{scanner: scanner, current: current, prev: prev}
}

func (p *Parser) expression() expr.Expr {
	return p.equality()
}

func (p *Parser) equality() expr.Expr {
	exp := p.comparison()
	for p.match(scanner.BangEqual, scanner.EqualEqual) {
		operator := p.previous()
		right := p.comparison()
		exp = expr.NewBinary(exp, operator, right)
	}
	return exp
}

func (p *Parser) comparison() expr.Expr {
	var exp = p.term()
	for p.match(scanner.Greater, scanner.GreaterEqual, scanner.Less, scanner.LessEqual) {
		operator := p.previous()
		right := p.term()
		exp = expr.NewBinary(exp, operator, right)
	}
	return exp
}

func (p *Parser) term() expr.Expr {
	var exp = p.factor()
	for p.match(scanner.Minus, scanner.Plus) {
		operator := p.previous()
		right := p.factor()
		exp = expr.NewBinary(exp, operator, right)
	}
	return exp
}

func (p *Parser) factor() expr.Expr {
	var exp = p.unary()
	for p.match(scanner.Star, scanner.Slash) {
		operator := p.previous()
		right := p.unary()
		exp = expr.NewBinary(exp, operator, right)
	}
	return exp
}

func (p *Parser) unary() expr.Expr {
	if p.match(scanner.Minus, scanner.Bang) {
		operator := p.previous()
		right := p.unary()
		return expr.NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() expr.Expr {
	switch {
	case p.match(scanner.False):
		return expr.NewBoolLiteral(false)
	case p.match(scanner.True):
		return expr.NewBoolLiteral(true)
	case p.match(scanner.Nil):
		return expr.NewNilLiteral()
	case p.match(scanner.Number):
		return expr.NewNumberLiteral(p.previous().Literal.Number)
	case p.match(scanner.String):
		return expr.NewStringLiteral(p.previous().Literal.String)
	case p.match(scanner.LeftParen):
		exp := p.expression()
		p.consume(scanner.RightParen, "Expected ')' after expression")
		return expr.NewGrouping(exp)
	}
}

func (p *Parser) consume(paren scanner.TokenType, s string) {

}
