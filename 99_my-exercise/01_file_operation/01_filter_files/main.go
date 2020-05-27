package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "os"
	"io/ioutil"
)

func main() {
	ss, err := getFileFullNames(`F:\download`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range ss {
		fmt.Println(s)
	}
}

func getFileFullNames(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	var s []string
	if err != nil {
		fmt.Println(err)
		return s, err
	}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			getFileFullNames(fileInfo.Name())
		} else {
			s = append(s, fileInfo.Name())
		}
	}
	return s, nil
}
