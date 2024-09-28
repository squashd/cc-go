package main

import (
	"fmt"
	"strings"
)

type (
	Node interface {
		TokenLiteral() string
		String() string
	}

	JSON struct {
		Value Node
	}
)

func (j *JSON) TokenLiteral() string {
	if j.Value != nil {
		return j.Value.TokenLiteral()
	} else {
		return ""
	}
}

func (j *JSON) String() string {
	if j.Value != nil {
		return j.Value.String()
	} else {
		return ""
	}
}

type Object struct {
	Pairs map[string]Node
}

func (o *Object) TokenLiteral() string {
	return "{"
}
func (o *Object) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for k, v := range o.Pairs {
		sb.WriteString(`"`)
		sb.WriteString(k)
		sb.WriteString(`"`)
		sb.WriteString(": ")
		sb.WriteString(v.String())
		sb.WriteString(", ")
	}
	str := sb.String()
	str = strings.TrimSuffix(str, ", ")
	str += "}"
	return str
}

type Array struct {
	Elements []Node
}

func (a *Array) TokenLiteral() string {
	return "["
}
func (a *Array) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, e := range a.Elements {
		sb.WriteString(e.String())
		if i < len(a.Elements)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

type Literal struct {
	Value any
	Type  TokenType
}

func (l *Literal) TokenLiteral() string {
	switch v := l.Value.(type) {
	case string:
		return "\"" + v + "\""
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case nil:
		return "null"
	default:
		return "unknown"
	}
}

func (l *Literal) String() string {
	return l.TokenLiteral()
}
