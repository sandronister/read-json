package db

import (
	"os"
)

func Load(path string) []byte {
	jsonFile, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return jsonFile

}
