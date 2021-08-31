package main

import "reflect"

type stack []interface{}

func (s stack) push(v interface{}) {
	s = append(s, v)
}

func main() {
	var c *interface{} = new(interface{})
	println(*c)
	*c = "cccc"
	// str, _ :=
	print("*c.type %s", reflect.TypeOf(c))
	println((*c).(string))
}
