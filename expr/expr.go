package expr

import "github.com/vlad-tokarev/glox/scanner"

type LitType int

type Visitor interface {
	VisitLiteral(node Literal) (w Visitor)
	VisitBinary(node Binary) (w Visitor)
	VisitUnary(node Unary) (w Visitor)
	VisitGrouping(node Grouping) (w Visitor)
}

type Expr interface {
	accept(v Visitor)
}

const (
	LitNumber LitType = iota
	LitString
	LitBool
	LitNil
)

type Literal struct {
	Number  float64
	String  string
	Bool    bool
	LitType LitType
}

func (l Literal) accept(v Visitor) {
	v.VisitLiteral(l)
}

func NewBoolLiteral(b bool) Literal {
	return Literal{Bool: b, LitType: LitBool}
}

func NewStringLiteral(s string) Literal {
	return Literal{String: s, LitType: LitString}
}

func NewNumberLiteral(n float64) Literal {
	return Literal{Number: n, LitType: LitNumber}
}

func NewNilLiteral() Literal {
	return Literal{LitType: LitNil}
}

type Binary struct {
	left     Expr
	operator scanner.Token
	right    Expr
}

func (b Binary) accept(v Visitor) {
	v.VisitBinary(b)
}

func NewBinary(left Expr, operator scanner.Token, right Expr) *Binary {
	return &Binary{left: left, operator: operator, right: right}
}

type Unary struct {
	operator scanner.Token
	right    Expr
}

func (u Unary) accept(v Visitor) {
	v.VisitUnary(u)
}

func NewUnary(operator scanner.Token, right Expr) Unary {
	return Unary{operator: operator, right: right}
}

type Grouping struct {
	expr Expr
}

func (g Grouping) accept(v Visitor) {
	v.VisitGrouping(g)
}

func NewGrouping(expr Expr) Grouping {
	return Grouping{expr: expr}
}
