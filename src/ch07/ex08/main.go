package main

import (
	"sort"
	"fmt"
)

var title = func(x, y *Track) bool {
	return x.Title < y.Title
}
var year = func(x, y *Track) bool {
	return x.Year < y.Year
}

func main() {
	fmt.Println("----------tracks data-----------")
	printTracks(tracksData)
	fmt.Println("----------main sort data-----------")
	mainSort()
	fmt.Println("----------stable sort data-----------")
	stableSort()
}

func mainSort() {
	data := createData()
	sortTable := MainSort{t: data}
	sortTable.addSortKey(title)
	sortTable.addSortKey(year)
	sort.Sort(&sortTable)
	printTracks(data)
}

func stableSort() {
	data := createData()
	sortTable := StableSort{t: data}
	sortTable.addSortKey(title)
	sortTable.addSortKey(year)
	for sortTable.HasNext() {
		sort.Stable(&sortTable)
	}
	printTracks(data)
}


func createData() []*Track {
	data := make([]*Track, len(tracksData))
	copy(data, tracksData)
	return data
}
