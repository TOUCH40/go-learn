package max_heap

// type Interface interface{
// 	sort.Interface
// 	Push(x interface{})// add x as element Len()
// 	Pop()interface{}// remove and return element len() -1
// }

// // Package sort provides primitives for sorting slices and user-defined
// // collections.
// package sort

// // A type, typically a collection, that satisfies sort.Interface can be
// // sorted by the routines in this package. The methods require that the
// // elements of the collection be enumerated by an integer index.
// type Interface interface {
// 	// Len is the number of elements in the collection.
// 	Len() int
// 	// Less reports whether the element with
// 	// index i should sort before the element with index j.
// 	Less(i, j int) bool
// 	// Swap swaps the elements with indexes i and j.
// 	Swap(i, j int)
// }

type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            {}
func (h hp) Less() int           {}
func (h hp) Swap() int           {}
func (h *hp) Push(v interface{}) {}
func (h *hp) Pop() interface{}   {}

func main() {
pos1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j && j == 3 {
				goto pos1

			}
		}
	}
}
