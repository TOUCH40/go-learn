// 分组
// 支持分组嵌套
// middleware//，作用在主分组上的中间件也会作用在子分组
// 子分组还应该带有自己的中间件
// 提供强大的扩展能力
// 还可以使用装饰器，更为灵活，效率未知
package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello geektutu\n")
	})
	// index of range for testing Recovery()
	r.GET("/recovery", func(c *gee.Context) {
		name := []string{"geektutu"}
		c.String(http.StatusOK, name[1000])
	})

	r.Run(":8000")
}
