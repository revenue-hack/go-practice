package ex02

import (
	"bytes"
	"sort"
	"strconv"
)

type MapIntSet struct {
	mapSet map[int]bool
}

func (m *MapIntSet) Has(x int) bool {
	if m == nil {
		return false
	}
	return m.mapSet[x]
}

func (m *MapIntSet) Add(x int) {
	if m == nil || m.mapSet == nil {
		m.mapSet = make(map[int]bool)
	}
	m.mapSet[x] = true
}

func (m *MapIntSet) UnionWith(t *MapIntSet) {
	if m == nil || m.mapSet == nil {
		m.mapSet = make(map[int]bool)
	}
	for i, tb := range t.mapSet {
		m.mapSet[i] = tb
	}
}

func (m *MapIntSet) String() string {
	if m == nil {
		return "{ }"
	}
	var keys []int
	for i, mb := range m.mapSet {
		if mb {
			keys = append(keys, i)
		}
	}
	sort.Ints(keys)
	keyLenIndex := len(keys) - 1
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, key := range keys {
		buf.WriteString(strconv.Itoa(key))
		if i != keyLenIndex {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}
