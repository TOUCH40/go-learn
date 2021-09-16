package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(k, value string) {
	c.Writer.Header().Add(k, value)
}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// 参数
	Params map[string]string
	// response info
	StatusCode int
	handlers   []HandlerFunc
	index      int
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// return josn
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Fail(code int, err string) {
	c.JSON(code, H{"message": err})
}

// 很牛啊，看半天才看懂
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	// 若前面的中间件未调用next方法，帮助它调用后面的中间件内容。
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}
