// os库 : 文件\文件夹创建,读取,移动,复制

// io库 : 文件内容的写入,修改,拼接

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 读取文件内容
	file, _ := os.OpenFile("./demo_unicode.html", 2, 0666)
	fileByte1, _ := ioutil.ReadAll(file)

	// 读取文件内容，更简易的方法
	fileByte2, _ := ioutil.ReadFile("./demo_unicode.html")

	// byte转string
	fileString := string(fileByte1)
	fileString2 := string(fileByte2)
	log.Println(fileString)
	log.Println(fileString2)
}
