package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface{
	// TokenLiteral returns the literal value of the token
	TokenLiteral() string

	String() string
}

// Statement is an interface that represents a statement node in the AST
type Statement interface {
	Node
	// statementNode is a marker method to distinguish statement nodes
	statementNode()
}
// Expression is an interface that represents an expression node in the AST
type Expression interface {
	Node
	// expressionNode is a marker method to distinguish expression nodes
	expressionNode()
}


type Program struct{
	Statements []Statement
}

func(p *Program)TokenLiteral() string{
	if len(p.Statements)>0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

type LetStatement struct{
	Token token.Token // The token.Let token
	Name *Identifier // The variable name
	Value Expression // The value to assign to the variable

}

func(ls *LetStatement)statementNode(){}
func(ls *LetStatement)TokenLiteral() string{
	return ls.Token.Literal
}

type Identifier struct{
	Token token.Token // The token.IDENT token
	Value string // The name of the variable
}

func(i *Identifier)expressionNode(){}
func(i *Identifier)TokenLiteral() string{
	return i.Token.Literal
}

type ReturnStatement struct{
	Token token.Token // The "return" token
	ReturnValue Expression // The value to return
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string{
	return rs.Token.Literal
}

func (p *Program) String() string{
	var out bytes.Buffer
	for _, s:= range p.Statements{
		out.WriteString(s.String())
	}

	return out.String()
}

func (i *Identifier) String() string{
	return i.Value
}
type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}

func(es *ExpressionStatement) statementNode() {}
func(es *ExpressionStatement) TokenLiteral() string{
	return es.Token.Literal
}

func (ls *LetStatement) String() string{
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil{
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string{
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil{
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string{
	if es.Expression != nil{
		return es.Expression.String()
	}

	return ""
}
	