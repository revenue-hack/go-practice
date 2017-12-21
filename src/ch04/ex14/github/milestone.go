package github

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

const milestoneUrl = "https://api.github.com/repos/revenue-hack/go-practice/milestones"

type Milestone struct {
	Number int
	Title string
	Description string
}

type MilestoneList struct {
	Milestones []*Milestone
}

func MilestonesRequest() MilestoneList {
	response, err := get(milestoneUrl)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("MileStonesRequest response panic")
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("read response status not 200")
		os.Exit(1)
	}
	var milestoneList MilestoneList
	if err := json.NewDecoder(response.Body).Decode(&(milestoneList.Milestones)); err != nil {
		fmt.Println("read response decode error")
		panic(err)
	}
	return milestoneList
}
