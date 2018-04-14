package ex09

import (
	"reflect"
	"testing"
)

func TestDecoder_Token(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected []Token
	}{
		{"(token)", []Token{StartList{}, Symbol{"token"}, EndList{}}},
		{`("token" token2)`, []Token{
			StartList{},
			String{"token"},
			Symbol{"token2"},
			EndList{}}},
		{"()", []Token{StartList{}, EndList{}}},
	} {
		dec := NewDecoder([]byte(c.input))
		for _, ex := range c.expected {
			result, err := dec.Token()
			if err != nil {
				t.Errorf("token error %v\n", err)
			}
			if !reflect.DeepEqual(result, ex) {
				t.Errorf("result = %v, expected = %v\n", result, ex)
			}
		}
	}
}
