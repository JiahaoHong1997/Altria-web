package main

import (
	"altria"
	"net/http"
)

func main() {
	r := altria.Saber()
	r.GET("/index", func(c *altria.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1") //	/v1
	{
		v1.GET("/", func(c *altria.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Altria!</h1>")
		})

		v1.GET("/hello", func(c *altria.Context) {
			// expect /hello?name=jiahaohong
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := v1.Group("/v2") //  /v1/v2
	{
		v2.GET("/hello/:name", func(c *altria.Context) {
			// expect /hello/jiahaohong
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *altria.Context) {
			c.JSON(http.StatusOK, altria.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
