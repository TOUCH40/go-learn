package main

import (
	// "bufio"
	"fmt"

	// "io"
	// "os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Path string `yaml:"path"`
}

func main() {
	configFile, err := ioutil.ReadFile(`conf.yaml`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	conf := new(config)
	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ss, err := getFileFullNames(conf.Path)
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
			s2, err2 := getFileFullNames(dirPath + `\` + fileInfo.Name())
			if err2 != nil {
				fmt.Println(err2.Error())
			} else {
				s = append(s, s2...)
			}

		} else {
			s = append(s, dirPath+`\`+fileInfo.Name())
		}
	}
	return s, nil
}
