package main

import (
	"altria"
	"net/http"
)

func main() {
	r := altria.Saber()
	r.GET("/", func(c *altria.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Altria!</h1>")
	})

	r.GET("/hello", func(c *altria.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *altria.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *altria.Context) {
		c.JSON(http.StatusOK, altria.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
