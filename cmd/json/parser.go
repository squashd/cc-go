package main

import "fmt"

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseNode() (Node, error) {
	var node Node
	var err error
	switch p.curToken.Type {
	case LBRACE:
		node, err = p.parseObject()
		if err != nil {
			return nil, err
		}
		return node, nil
	case LBRACKET:
		node, err = p.parseArray()
		if err != nil {
			return nil, err
		}
		return node, nil
	case STRING, INT, FLOAT, TRUE, FALSE, NULL:
		node, err = p.parseLiteral()
		if err != nil {
			return nil, err
		}
		return node, nil
	default:
		return nil, fmt.Errorf("Invalid token: %s", p.curToken.Type)
	}
}

func (p *Parser) InvalidJSONErr() error {
	return fmt.Errorf("invalid JSON — token type: %s, token liter: %s, read position: %d", p.curToken.Type, p.curToken.Literal, p.l.readPosition)
}

func (p *Parser) ParseJSON() (*JSON, error) {
	var node Node
	var err error

	switch p.curToken.Type {
	case LBRACE:
		node, err = p.parseObject()
		if err != nil {
			return nil, err
		}
	case LBRACKET:
		node, err = p.parseArray()
		if err != nil {
			return nil, err
		}
	case STRING, INT, FLOAT, TRUE, FALSE, NULL:
		node, err = p.parseLiteral()
		if err != nil {
			return nil, err
		}
	default:
		return nil, p.InvalidJSONErr()
	}

	return &JSON{Value: node}, nil
}

func (p *Parser) parseObject() (*Object, error) {
	obj := &Object{Pairs: make(map[string]Node)}

	if p.curToken.Type != LBRACE {
		return nil, p.InvalidJSONErr()
	}

	p.nextToken()

	for p.curToken.Type != RBRACE && p.curToken.Type != EOF {
		if p.curToken.Type != STRING {
			return nil, p.InvalidJSONErr()
		}
		key := p.curToken.Literal

		p.nextToken()
		if p.curToken.Type != COLON {
			return nil, p.InvalidJSONErr()
		}

		p.nextToken()
		value, err := p.parseNode()
		if err != nil {
			return nil, p.InvalidJSONErr()
		}

		obj.Pairs[key] = value

		p.nextToken()
		if p.curToken.Type == COMMA {
			p.nextToken()
		}
	}

	return obj, nil
}

func (p *Parser) parseArray() (*Array, error) {
	var value Node
	var err error
	arr := &Array{}

	if p.curToken.Type != LBRACKET {
		return nil, p.InvalidJSONErr()
	}

	p.nextToken()

	for p.curToken.Type != RBRACKET && p.curToken.Type != EOF {
		value, err = p.parseNode()
		if err != nil {
			return nil, p.InvalidJSONErr()
		}
		arr.Elements = append(arr.Elements, value)

		p.nextToken()
		if p.curToken.Type == COMMA {
			p.nextToken()
		}
	}

	return arr, nil
}

func (p *Parser) parseLiteral() (*Literal, error) {
	switch p.curToken.Type {
	case STRING, INT, FLOAT:
		return &Literal{Value: p.curToken.Literal}, nil
	case TRUE:
		return &Literal{Value: true}, nil
	case FALSE:
		return &Literal{Value: false}, nil
	case NULL:
		return &Literal{Value: nil}, nil
	}
	return nil, p.InvalidJSONErr()
}
