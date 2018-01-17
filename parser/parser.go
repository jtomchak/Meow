package parser

import (
	"meow/ast"
	"meow/lexer"
	"meow/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//Read two tokens, so curToken and peekToken are both setup
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	//assign next to current and advance with NextToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
