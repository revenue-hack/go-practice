package github

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

func envLoad() Credential {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error load .env file")
	}
	return Credential{os.Getenv("USERNAME"), os.Getenv("PASSWORD")}
}