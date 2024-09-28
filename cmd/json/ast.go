package main

type (
	Node interface {
		TokenLiteral() string
	}

	Statement interface {
		Node
		statementNode()
	}

	Expression interface {
		Node
		expressionNode()
	}

	Program struct {
		Satements []Statement
	}
)

func (p *Program) TokenLiteral() string {
	if len(p.Satements) > 0 {
		return p.Satements[0].TokenLiteral()
	} else {
		return ""
	}
}
