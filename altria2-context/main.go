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
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *altria.Context) {
		c.JSON(http.StatusOK, altria.H{
			"username":	c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
