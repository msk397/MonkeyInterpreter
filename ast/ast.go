package ast

import (
	"MonkeyInterpreter/token"
	"bytes"
)

type Node interface {
	//该方法用于返回语法单元的值，比如关键字的值、标识符的值、运算符的值等
	TokenLiteral() string
	String() string
}

// Statement
// 语句
type Statement interface {
	Node
	statementNode()
}

// Expression
// 表达式
type Expression interface {
	Node
	expressionNode()
}

// Program AST的根节点
type Program struct {
	Statements []Statement
}

// TokenLiteral
// 返回语法单元的值
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String
// 用于实现Node接口
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement
// let <标识符> = <表达式>
type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

// statementNode
// 用于实现Statement接口
func (ls *LetStatement) statementNode() {}

// TokenLiteral
// 用于实现Node接口
// 返回let关键字的值
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String
// 用于实现Node接口
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier
type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string
}

// expressionNode
// let语句中的标识符不会出现表达式，但还是要实现Expression接口，因为其他地方的标识符可能会出现表达式
func (i *Identifier) expressionNode() {}

// String
func (i *Identifier) String() string { return i.Value }

// TokenLiteral
// 用于实现Node接口
// 返回标识符的值
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement
// return 语句
// return <表达式>
type ReturnStatement struct {
	Token       token.Token // token.RETURN 词法单元
	ReturnValue Expression  // 表达式
}

// statementNode
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral
// 返回return关键字的值
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement
// 表达式语句
// <表达式>
type ExpressionStatement struct {
	Token      token.Token // 表达式的第一个词法单元
	Expression Expression  // 表达式
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
