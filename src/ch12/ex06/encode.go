package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		log.Println(v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if isZero(v.Field(i)) {
				continue
			}
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "%s", "t")
		} else {
			fmt.Fprintf(buf, "%v", nil)
		}

	case reflect.Interface:
		buf.WriteByte('(')
		fmt.Fprintf(buf, "%s ", v.Elem().Type().String())
		if err := encode(buf, v.Elem()); err != nil {
			log.Fatal(err)
		}
		buf.WriteByte(')')

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%.2f %.2f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func main() {
	type zero struct {
	}
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
		zeroMap         zero
	}
	var zeroMap map[string]string
	var zeroComp64 complex64
	var zeroStruct zero
	strangelove := Movie{
		Title:       "Dr. Strangelove",
		Subtitle:    "How I Learned to Stop Worrying and Love the Bomb",
		Year:        0,
		Color:       false,
		HogeComp64:  zeroComp64,
		HogeComp128: complex(2, 4),
		HogeBool:    true,
		Actor:       zeroMap,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Inter:   []int{1, 2, 3},
		zeroMap: zeroStruct,
	}
	result, err := Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(result))
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Interface, reflect.Ptr, reflect.Struct:
		return v.IsValid()
	}
	return false
}
