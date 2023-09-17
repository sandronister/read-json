package main

import (
	"fmt"

	user_repository "lab.message/repository/user"
)

func main() {
	user_repository.Init()
	user := user_repository.Search("Elliot")
	fmt.Print(user)
}
