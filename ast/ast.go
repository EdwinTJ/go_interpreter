package ast

import "monkey/token"

type Node interface{
	// TokenLiteral returns the literal value of the token
	TokenLiteral() string
}

// Statement is an interface that represents a statement node in the AST
type Statements interface {
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
	Statements []Statements
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