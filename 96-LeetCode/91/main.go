package main

import("fmt")

func main()  {
	fmt.Println(numDecodings("10"))
}

func numDecodings( s string)int  {
	n:=len(s)
	var c  [5]int
	fmt.Println(c[0:3])
	return len(c)+n
}