package ex09

import (
	"bytes"
	"fmt"
	"strconv"
	"text/scanner"
)

type Symbol struct {
	Value string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct {
}

type EndList struct {
}

type Token interface {
}

type Decoder struct {
	lex *lexer
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func NewDecoder(data []byte) *Decoder {
	var dec Decoder
	dec.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec.lex.scan.Init(bytes.NewReader(data))
	dec.lex.next()
	return &dec
}

func (dec *Decoder) Token() (Token, error) {
	switch dec.lex.token {
	case scanner.Ident:
		symbol := dec.lex.text()
		dec.lex.next()
		return Symbol{symbol}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(dec.lex.text())
		dec.lex.next()
		return Int{i}, nil
	case scanner.String:
		s, _ := strconv.Unquote(dec.lex.text())
		dec.lex.next()
		return String{s}, nil
	case '(':
		dec.lex.next()
		return StartList{}, nil
	case ')':
		dec.lex.next()
		return EndList{}, nil
	}
	return nil, fmt.Errorf("error token = %v\n", dec.lex.token)
}
