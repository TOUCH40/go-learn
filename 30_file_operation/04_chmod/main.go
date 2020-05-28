package main

import (
	"os"
)

func main() {
	os.Chmod("aa/testA", 0666)
}
