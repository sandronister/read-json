package user

import (
	"encoding/json"

	"lab.message/db"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type UserRepository struct {
	list Users
}

var list Users

func Init() {
	jsonFile := db.Load("./files/user.json")
	json.Unmarshal(jsonFile, &list)
}

func Search(name string) User {
	var found User
	for _, user := range list.Users {
		if user.Name == name {
			found = user
			break
		}
	}

	return found
}
