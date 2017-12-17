package main

import(
	"gopl.io/ch4/github"
	"os"
	"fmt"
	"time"
)
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	var months, years, others []*github.Issue
	for _, item := range result.Items {
		if isLessThanOneMonth(item.CreatedAt) {
			months = append(months, item)
		} else if isLessThanOneYear(item.CreatedAt) {
			years = append(years, item)
		} else {
			others = append(others, item)
		}
	}
	display(months, years, others)
}

func display(months, years, others []*github.Issue) {
	fmt.Println("1ヶ月未満")
	for _, item := range months {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("1年未満")
	for _, item := range years {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("その他")
	for _, item := range others {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func isLessThanOneMonth(compareTime time.Time) bool {
	return compareTime.After(time.Now().AddDate(0, -1, 0))
}

func isLessThanOneYear(compareTime time.Time) bool {
	return compareTime.After(time.Now().AddDate(-1, 0, 0))
}

