package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close() // 主协程结束时，关闭listener

	// 接收多个用户
	for {
		//fmt.Println("服务器等待客户端建立连接...")
		// 等待客户端的连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("accept err:", err)
			return
		}
		// defer conn.Close() // 使用结束，断开与客户端链接
		// 新建一个携程
		go HandleConn(conn)
		// 接收客户端数据
		// buf := make([]byte, 1024) // 创建1024大小的缓冲区，用于read
		// n, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("read err:", err)
		// 	return
		// }
	}

	//fmt.Println("客户端与服务器连接建立成功...")

	// fmt.Println("服务器读到:", string(buf[:n])) // 读多少，打印多少
}
