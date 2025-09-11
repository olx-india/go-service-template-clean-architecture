package integrationtests

import (
	"encoding/json"
	"os"
)

func LoadJSON(path string) interface{} {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	var v interface{}
	if err := json.NewDecoder(f).Decode(&v); err != nil {
		panic(err)
	}
	return v
}
