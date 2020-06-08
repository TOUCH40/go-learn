package main

import "fmt"

func main() {
	In := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"
	Out := findLadders(beginWord, endWord, In)
	fmt.Println(Out)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	mapWordList := make(map[string]bool)
	retSlice := make([][]string, 0)

	for _, word := range wordList {
		mapWordList[word] = true
	}

	_, okB := mapWordList[beginWord]
	_, okE := mapWordList[endWord]

	if okB && okE {

	}

	return retSlice
}
