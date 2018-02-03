package eval

import (
	"fmt"
	"strconv"
)

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'g', -2, 64)
}

func (v Var) String() string {
	return string(v)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var str = fmt.Sprintf("%s(", c.fn)
	for i, arg := range c.args {
		if i != 0 {
			str += ", "
		}
		str += arg.String()
	}
	str += ")"
	return str
}
