package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := new(person)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
