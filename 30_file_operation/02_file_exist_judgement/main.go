package main

import (
	"io"
	"os"
)

func main() {
	os.MkdirAll("aa", 0777)

	// 创建了文件 testA, 并且写入了内容
	testA, _ := os.Create("aa/testA")
	io.WriteString(testA, "the teatA")
	// 如果需要安全判断, 可以使用 os.Stat 配合 os.IsNotExist
	if _, err := os.Stat("aa/testA"); os.IsNotExist(err) {
		// 当文件不存在, 才执行创建
		os.Create("aa/testA")
	}
	os.Chmod("aa/testA", 0777)
}
