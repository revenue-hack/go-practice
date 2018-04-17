package cycle

import "testing"

func TestIsCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	d := &link{value: "d"}
	d.tail = nil
	for _, test := range []struct {
		linktype *link
		expected bool
	}{
		{a, true},
		{b, true},
		{c, true},
		{d, false},
	} {
		if IsCycle(test.linktype) != test.expected {
			t.Errorf("is cycle fail input = %v expeced = %v\n", test.linktype, test.expected)
		}
	}
}
