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
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintln(buf, "")
				fmt.Fprint(buf, "\t")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i != v.NumField()-1 {
				fmt.Fprintln(buf, "")
			}
		}
		buf.WriteByte(')')

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "%s", "t")
		} else {
			fmt.Fprintf(buf, "%v", nil)
		}

	case reflect.Interface:
		log.Println(v.Elem())
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
				fmt.Fprint(buf, "\t")
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
			if i != len(v.MapKeys())-1 {
				fmt.Fprintln(buf, "")
			}
		}
		buf.WriteByte(')')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func main() {
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
	result, err := Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}
