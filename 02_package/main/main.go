package main

// 在我们引入第三方的时候，会在三个地方进行查找
// 1.GOROOT路径
// 2.GOPATH路径
// 3.在原目录中的vendor目录下进行查找

import (
	"fmt"
	// "github.com/GoesToEleven/GolangTraining/02_package/stringutil"
	// "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	winniepooh "github.com/touch40/go-learn/02_package/incomefromalaska"
	"github.com/touch40/go-learn/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
}
