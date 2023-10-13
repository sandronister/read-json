package main

import (
	"fmt"

	"lab.message/repository/user"
)

func main() {
	user.Init()
	founded := user.Search("Elliot")
	fmt.Print(founded)
}
