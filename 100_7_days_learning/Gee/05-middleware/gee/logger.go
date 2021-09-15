package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Caculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.req.RequestURI, time.since(t))
	}
}
