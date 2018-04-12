package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"text/scanner"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	HogeBool        bool
	HogeComp64      complex64
	HogeComp128     complex128
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Inter           interface{}
}

func main() {
	strangelove := Movie{
		Title:       "Dr. Strangelove",
		Subtitle:    "How I Learned to Stop Worrying and Love the Bomb",
		Year:        1964,
		Color:       false,
		HogeComp64:  complex(1, 2),
		HogeComp128: complex(2, 4),
		HogeBool:    true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Inter: []int{1, 2, 3},
	}
	data, err := Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	var movie Movie
	dec := &Decoder{bytes.NewReader(data)}
	if err = dec.Decode(&movie); err != nil {
		log.Fatal(err)
	}
	log.Println(movie)
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
	read(lex, reflect.ValueOf(out).Elem())
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

func read(lex *lexer, v reflect.Value) {
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
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
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
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		if lex.text() == "\"[]int\"" {
			var tmp []int
			val := reflect.New(reflect.TypeOf(tmp)).Elem()
			lex.next()
			read(lex, val)
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
