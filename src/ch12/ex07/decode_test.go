package main

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testMovie := Movie{
		Title:       "Dr. Strangelove",
		Subtitle:    "How I Learned to Stop Worrying and Love the Bomb",
		Year:        1964,
		Color:       false,
		HogeComp64:  complex(1, 2),
		HogeComp128: complex(2, 4),
		Actor: map[string]string{
			"Dr. Strangelove":            "",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		HogeBool: true,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Inter: []int{1, 2, 3},
	}
	data, err := Marshal(testMovie)
	if err != nil {
		t.Error("marshal error")
	}
	var expected Movie
	err = Unmarshal(data, &expected)
	if err != nil {
		t.Error("unmarshal error")
	}
	if !reflect.DeepEqual(expected, testMovie) {
		t.Error("test unmarshal error")
	}
}
