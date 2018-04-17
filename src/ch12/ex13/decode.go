package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

type Movie struct {
	Title       string            `sexpr:"title"`
	Subtitle    string            `sexpr:"subtitle"`
	Year        int               `sexpr:"year"`
	Color       bool              `sexpr:"color"`
	HogeBool    bool              `sexpr:"hogebool"`
	HogeComp64  complex64         `sexpr:"hogecomp64"`
	HogeComp128 complex128        `sexpr:"hogecomp128"`
	Actor       map[string]string `sexpr:"actor"`
	Oscars      []string          `sexpr:"oscars"`
	Sequel      *string           `sexpr:"sequel"`
	Inter       interface{}       `sexpr:"inter"`
}

type Decoder struct {
	reader io.Reader
}

func (dec *Decoder) Decode(out interface{}) error {
	buf := new(bytes.Buffer)
	io.Copy(buf, dec.reader)
	return Unmarshal(buf.Bytes(), out)
}

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	fields := make(map[string]string)
	v := reflect.ValueOf(out).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("sexpr")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Type().Field(i).Name
	}
	read(lex, reflect.ValueOf(out).Elem(), fields)
	return nil
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

func read(lex *lexer, v reflect.Value, fields map[string]string) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		if lex.text() == "\"t\"" {
			v.SetBool(true)
		} else {
			s, _ := strconv.Unquote(lex.text())
			v.SetString(s)
		}
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '#':
		lex.next()
		lex.next()
		lex.next()
		real := lex.text()
		pReal, err := strconv.ParseFloat(real, 64)
		if err != nil {
			log.Fatal(err)
		}
		lex.next()
		imag := lex.text()
		pImag, err := strconv.ParseFloat(imag, 64)
		if err != nil {
			log.Fatal(err)
		}
		v.SetComplex(complex(pReal, pImag))
		lex.next()
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v, fields)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value, fields map[string]string) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i), fields)
		}

	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item, fields)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(fields[name]), fields)
			lex.consume(')')
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key, fields)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value, fields)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		if lex.text() == "\"[]int\"" {
			var tmp []int
			val := reflect.New(reflect.TypeOf(tmp)).Elem()
			lex.next()
			read(lex, val, fields)
			v.Set(val)
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
