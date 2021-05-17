package main

import (
	"fmt"
)

func main() {

	c := memLeak(8, 11)
	fmt.Println(c)
}

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	cd := [n][m]byte{}
	for i := range box {
		for j := range box[i] {
			cd[n-j][i] = box[i][j]
		}
	}
	return cd
}
