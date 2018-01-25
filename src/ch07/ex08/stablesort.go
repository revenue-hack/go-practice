package main

type StableSort struct {
	t []*Track
	lesses []func(x, y *Track) bool
	lessIndex int
}

func (s *StableSort) addSortKey(less func(x, y *Track) bool) {
	s.lesses = append(s.lesses, less)
	s.lessIndex++
}

func (s *StableSort) Len() int {
	return len(s.t)
}

func (s *StableSort) Less(i, j int) bool {
	return s.lesses[s.lessIndex](s.t[i], s.t[j])
}

func (s *StableSort) Swap(i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func (s *StableSort) HasNext() bool {
	s.lessIndex--
	return s.lessIndex >= 0
}
