package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		{"{10, x, -y}", Env{"x": 10, "y": -20}, "10"},
		{"{10, x, -y}", Env{"x": 12, "y": 6}, "-6"},
		{"{10, x, -y}", Env{"x": -20, "y": 6}, "-20"},
		//!+Eval
	}
	for _, test := range tests {

		expr, err := Parse(test.expr)
		if err != nil {
			t.Error("parse error")
			t.Error(err)
			continue
		}
		reExpr, err := Parse(expr.String())
		if err != nil {
			t.Error("reparse error")
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", reExpr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("unexpected got: %s\twant: %s\n", got, test.want)
		}
	}
}
