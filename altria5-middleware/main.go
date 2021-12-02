package main

import (
	"altria"
	"log"
	"net/http"
	"time"
)

func onlyForV1() altria.HandlerFunc {
	return func(c *altria.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v1", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := altria.Saber()
	r.Use(altria.Logger()) // global middleware
	r.GET("/", func(c *altria.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(onlyForV1()) // v1 group middleware

	v1.GET("/hello/:name", func(c *altria.Context) {
		// expect /hello/jiahaohong
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Run(":9999")
}
