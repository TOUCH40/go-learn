package tcp

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
	}
}
