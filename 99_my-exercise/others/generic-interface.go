package main
 
import (
	"fmt"
)
 
type Lesser[T any] interface {
	Less(y T) bool
}
 
func Sort[Elem Lesser[Elem]](list []Elem) []Elem {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if !list[j].Less(list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
 
type Item struct {
	value string
}
 
func (x Item) Less(y Item) bool { return x.value < y.value }
 
type Int int
 
func (x Int) Less(y Int) bool { return x < y }
 
func main() {
	items := []Item{Item{"a"}, Item{"c"}, Item{"d"}, Item{"b"}}
	fmt.Println("result=", Sort(items))
 
	fmt.Println("result=", Sort([]Int{2, 3, 4, 1}))
}