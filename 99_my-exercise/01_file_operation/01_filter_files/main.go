package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type config struct {
	Path string `yaml:"path"`
}

var conf config

func main() {
	configFile, err := ioutil.ReadFile(`conf.yaml`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
			ism, err3 := isFileNameEndWith(fileInfo.Name(), "vep")
			if err3 != nil {
				fmt.Println(err3)
				continue
			}
			if ism {
				s = append(s, dirPath+`\`+fileInfo.Name())
			}

		}
	}
	return s, nil
}

func isFileNameEndWith(name string, extension string) (bool, error) {
	reg := fmt.Sprintf(`\.%s$`, extension) //`\.vep$`
	ret := regexp.MustCompile(reg)
	if ret == nil {
		fmt.Println("not regexp")
		return false, errors.New("not regexp")
	}
	isVep, err := regexp.MatchString(reg, name)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return isVep, nil
}

func gen(paths []string, target string) <-chan [2]string {
	out := make(chan [2]string)
	for _, path := range paths {
		changePath := [2]string{path, strings.Replace(path, conf.Path, target, 1)}
		out <- changePath
	}
	close(out)
	return out
}
