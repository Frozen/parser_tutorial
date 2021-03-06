package pg

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

//https://github.com/hoterran/sqlparser/blob/master/sql.y

//go:generate goyacc -l -o parser.go parser.y

// Parse parses the input and returns the result.
func Parse(input string) (Ast, error) {
	l := newLex([]byte(input))
	_ = yyParse(l)
	return l.result, l.err
}

var keywords = map[string]int{
	"let":   Let,
	"func":  FuncT,
	"match": Match,
	"case":  Case,
}

type lex struct {
	input  []byte
	pos    int
	result Ast
	err    error
}

func newLex(input []byte) *lex {
	return &lex{
		input: input,
	}
}

// Lex satisfies yyLexer.
func (l *lex) Lex(lval *yySymType) int {
	return l.scanNormal(lval)
}

func (l *lex) scanNormal(lval *yySymType) int {
	for b := l.next(); b != 0; b = l.next() {
		runeB := rune(b)
		_ = runeB
		switch {
		case b == '=':
			return l.scanEq(lval)
		case b == '(':
			return OpenB
		case b == ')':
			return CloseB
		case b == '{':
			return OpenF
		case b == '}':
			return CloseF
		case b == '[':
			return OpenS
		case b == ']':
			return CloseS
		case b == ',':
			return Comma
		case b == ':':
			return Colon
		case b == '%':
			return Mod
		case b == '+':
			return Plus
		case unicode.IsSpace(rune(b)):
			continue
		case b == '"':
			return l.scanString(lval)
		case unicode.IsDigit(rune(b)) || b == '+' || b == '-':
			l.backup()
			return l.scanNum(lval)
		case unicode.IsLetter(rune(b)) || b == '_':
			l.backup()
			return l.scanLiteral(lval)

		default:
			return int(b)
		}
	}
	return 0
}

var escape = map[byte]byte{
	'"':  '"',
	'\\': '\\',
	'/':  '/',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
}

func (l *lex) scanString(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for b := l.next(); b != 0; b = l.next() {
		switch b {
		case '\\':
			// TODO(sougou): handle \uxxxx construct.
			b2 := escape[l.next()]
			if b2 == 0 {
				return LexError
			}
			buf.WriteByte(b2)
		case '"':
			lval.str = buf.String()
			return String
		default:
			buf.WriteByte(b)
		}
	}
	return LexError
}

func (l *lex) scanNum(lval *yySymType) int {
	defer func() {
		fmt.Println("scanNum", lval)
	}()
	buf := bytes.NewBuffer(nil)
	for {
		b := l.next()
		switch {
		case unicode.IsDigit(rune(b)):
			buf.WriteByte(b)
		case strings.IndexByte(".+-eE", b) != -1:
			buf.WriteByte(b)
		default:
			l.backup()
			//val, err := strconv.ParseFloat(buf.String(), 64)
			//if err != nil {
			//	return LexError
			//}
			lval.str = buf.String()
			return Number
		}
	}
}

//var literal = map[string]interface{}{
//	"true":  true,
//	"false": false,
//	"null":  nil,
//}

func (l *lex) scanLiteral(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for {
		b := l.next()
		switch {
		case unicode.IsLetter(rune(b)):
			buf.WriteByte(b)
		case unicode.IsDigit(rune(b)):
			buf.WriteByte(b)
		case b == '_':
			buf.WriteByte(b)
		default:
			l.backup()
			val, ok := keywords[buf.String()]
			if !ok {
				lval.str = buf.String()
				return Literal
			}
			lval.val = val
			return val
		}
	}
}

func (l *lex) scanEq(lval *yySymType) int {
	next := l.next()
	switch {
	case next == '>':
		return RightArrow
	case next == '=':
		return Compare
	}
	l.backup()
	return Eq
}

func (l *lex) backup() {
	if l.pos == -1 {
		return
	}
	l.pos--
}

func (l *lex) next() byte {
	if l.pos >= len(l.input) || l.pos == -1 {
		l.pos = -1
		return 0
	}
	l.pos++
	return l.input[l.pos-1]
}

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	l.err = errors.New(s)
}
