package main

import (
	"io"
	"os"
)

func main() {
	// 设定工作路径, 一般在项目刚开始设定一次即可, 不要在异步过程中修改工作路径
	// 默认是程序执行的路径
	os.Chdir("./")

	// 创建文件夹
	os.MkdirAll("./aa/bb/c1", 0777)
	os.MkdirAll("./aa/bb/c2", 0777)

	// 创建文件
	os.Create("./aa/bb/c1/file.go")

	// 移动文件
	os.Rename("./aa/bb/c1/file.go", "./aa/bb/c2/file.go")

	// 打开文件,得到一个 *File 对象, 用于后续的写入
	file, _ := os.OpenFile("./aa/bb/c2/file.go", 2, 0666)

	// 写入文件内容
	io.WriteString(file, `
        package main
        func main(){
            log.Println("我是由程序写入的代码")
        }
    `)

	// 也可以直接调用file里的函数写入内容
	file.WriteString("add string")

	// 拷贝文件, 拷贝其实就是创建一个文件, 然后写入文件内容
	src1, _ := os.Create("./aa/bb/c1/file-copy1.go")
	io.Copy(file, src1) // 把文件file, 写入src1文件

	// 删除文件或文件夹
	os.Create("./aa/bb/c1/file-delete.go") // 创建一个文件用于删除
	os.RemoveAll("./aa/bb/c1/file-delete.go")

}
