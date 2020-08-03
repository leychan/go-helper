package go_helper

import "os"

func GetUser() string {
	return os.Getenv("USER")
}
