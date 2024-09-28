package main

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

func (p *Parser) parseNode() Node {
	switch p.curToken.Type {
	case LBRACE:
		return p.parseObject()
	case LBRACKET:
		return p.parseArray()
	case STRING, INT, FLOAT, TRUE, FALSE, NULL:
		return p.parseLiteral()
	default:
		return nil
	}
}

func (p *Parser) ParseJSON() *JSON {
	switch p.curToken.Type {
	case LBRACE:
		return &JSON{Value: p.parseObject()}
	case LBRACKET:
		return &JSON{Value: p.parseArray()}
	case STRING, INT, FLOAT, TRUE, FALSE, NULL:
		return &JSON{Value: p.parseLiteral()}
	default:
		return nil
	}
}

func (p *Parser) parseObject() *Object {
	obj := &Object{Pairs: make(map[string]Node)}

	if p.curToken.Type != LBRACE {
		return nil
	}

	p.nextToken()

	for p.curToken.Type != RBRACE && p.curToken.Type != EOF {
		if p.curToken.Type != STRING {
			return nil
		}
		key := p.curToken.Literal

		p.nextToken()
		if p.curToken.Type != COLON {
			return nil
		}

		p.nextToken()
		value := p.parseNode()

		obj.Pairs[key] = value

		p.nextToken()
		if p.curToken.Type == COMMA {
			p.nextToken()
		}
	}

	return obj
}

func (p *Parser) parseArray() *Array {
	arr := &Array{}

	if p.curToken.Type != LBRACKET {
		return nil
	}

	p.nextToken()

	for p.curToken.Type != RBRACKET && p.curToken.Type != EOF {
		value := p.parseNode()
		arr.Elements = append(arr.Elements, value)

		p.nextToken()
		if p.curToken.Type == COMMA {
			p.nextToken()
		}
	}

	return arr
}

func (p *Parser) parseLiteral() *Literal {
	switch p.curToken.Type {
	case STRING, INT, FLOAT:
		return &Literal{Value: p.curToken.Literal}
	case TRUE:
		return &Literal{Value: true}
	case FALSE:
		return &Literal{Value: false}
	case NULL:
		return &Literal{Value: nil}
	}
	return nil
}
