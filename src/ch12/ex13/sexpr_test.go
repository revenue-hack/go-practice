package sexpr

import (
	"bytes"
	"log"
	"reflect"
	"strings"
	"testing"
)

var testMovie = Movie{
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

func TestSexprDecode(t *testing.T) {
	data, err := Marshal(testMovie)
	if err != nil {
		log.Fatal(err)
	}
	var input Movie
	dec := &Decoder{bytes.NewReader(data)}
	if err = dec.Decode(&input); err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(input, testMovie) {
		t.Error("test decode error")
	}
}

var fields = []string{"Title", "Subtitle", "Year",
	"Color", "HogeBool", "HogeComp64", "HogeComp128",
	"Actor", "Oscars", "Sequel", "Inter"}

func TestSexprEncode(t *testing.T) {
	data, err := Marshal(testMovie)
	if err != nil {
		log.Fatal(err)
	}
	dataStr := string(data)
	for _, field := range fields {
		if strings.Contains(dataStr, "("+field) {
			t.Errorf("encode error %s\n", field)
		}
	}
}
