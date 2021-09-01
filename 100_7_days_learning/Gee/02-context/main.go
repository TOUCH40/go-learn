package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8000")
}

// 封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用
// 解析动态路由
// 中间件

// 封装前
// 	r.POST("/login", func(c *gee.Context) {
// 		c.JSON(http.StatusOK, gee.H{
// 			"username": c.PostForm("username"),
// 			"password": c.PostForm("password"),
// 		})
// 	})

// 	r.Run(":8000")
// }

// obj = map[string]interface{}{
// 	"name": "geektutu",
// 	"password": "1234",
// }
// w.Header().Set("Content-Type", "application/json")
// w.WriteHeader(http.StatusOK)
// encoder := json.NewEncoder(w)
// if err := encoder.Encode(obj); err != nil {
// 	http.Error(w, err.Error(), 500)
// }

// 封装后

// c.JSON(http.StatusOK, gee.H{
// 	"username": c.PostForm("username"),
// 	"password": c.PostForm("password"),
// })
