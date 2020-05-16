package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
/*
闭包帮助我们限制多个函数使用的变量的范围
如果没有闭包，两个或多个函数访问同一个变量时
该变量需要是包范围（全局）
*/
