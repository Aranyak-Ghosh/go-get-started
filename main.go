package main

import (
	"encoding/json"
	"fmt"
)

type sample struct {
	A int64 `json:"a"`
}

type SampleArray[T interface{}] []T

func (s *SampleArray[T]) UnmarshalJSON(data []byte) error {
	var out []T
	err := json.Unmarshal(data, &out)
	if err != nil {
		return err
	} else {
		*s = out
		return nil
	}
}

// Map, foreach, push, pop, removeAt, filter, 

func main() {
	response := []byte(`[
	{
		"a": 123
	}
	]`)

	var out SampleArray[sample]

	l := len(out)

	fmt.Println(l)

	err := json.Unmarshal(response, &out)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		dat, _ := json.Marshal(out[0])
		fmt.Printf(string(dat))
	}
}
