package main

import (
	"fmt"

	"lab.message/repository/user"
)

func main() {
	user.Init()
	user := user.Search("Elliot")
	fmt.Print(user)
}
