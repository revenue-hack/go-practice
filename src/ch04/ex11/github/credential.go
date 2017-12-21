package github

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func envLoad() Credential {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error load .env file")
	}
	return Credential{os.Getenv("USERNAME"), os.Getenv("PASSWORD")}
}
