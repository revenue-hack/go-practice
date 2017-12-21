package github

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

const userUrl = "https://api.github.com/users"

type User struct {
	Login string
	Id int
	Url string
}

type UserList struct {
	Users []*User
}

func UsersRequest() UserList {
	response, err := get(userUrl)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("usersRequest response panic")
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("read response status not 200")
		os.Exit(1)
	}
	var userList UserList
	if err := json.NewDecoder(response.Body).Decode(&(userList.Users)); err != nil {
		fmt.Println("read response decode error")
		panic(err)
	}
	return userList
}
