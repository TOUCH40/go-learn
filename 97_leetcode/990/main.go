package main

import (
	"fmt"
)

func main() {
	In := []string{"c==c", "b==d", "x!=z"}
	ret := equationsPossible(In)
	fmt.Println(ret)
}

func equationsPossible(equations []string) bool {
	Out := true
	equalMap := make(map[string]bool)
	for _, equation := range equations {
		if equalMap[1] == rune('=') {
			equalMap[equation] = true
		} else {
			equalMap[false] = true
		}
	}

	return Out
}
