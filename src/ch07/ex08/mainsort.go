package main

type MainSort struct {
	t []*Track
	LessFuncs
}
type LessFuncs struct {
	lesses []func(x, y *Track) bool
}

func (l *LessFuncs) addSortKey(less func(x, y *Track) bool) {
	l.lesses = append(l.lesses, less)
}

func (m *MainSort) Len() int {
	return len(m.t)
}

func (m *MainSort) Less(i, j int) bool {
	length := len(m.LessFuncs.lesses)
	var k int
	for k = 0; k < length - 1; k++ {
		switch {
		case m.LessFuncs.lesses[k](m.t[i], m.t[j]):
			return true
		case m.LessFuncs.lesses[k](m.t[j], m.t[i]):
			return false
		}
	}
	return m.LessFuncs.lesses[k](m.t[i], m.t[j])
}

func (m *MainSort) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

