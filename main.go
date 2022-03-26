package main

import (
	"encoding/json"
	"fmt"
)

type sample struct {
	A int64 `json:"a"`
}

func main() {
	response := []byte(`
	{
		"a": 123
	}`)

	var out sample
	err := json.Unmarshal(response, &out)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		dat, _ := json.Marshal(out)
		fmt.Printf(string(dat))
	}
}
