package main

import (
	"altria"
	"net/http"
)

func main() {
	r := altria.Saber()
	r.GET("/", func(c *altria.Context) {
		c.String(http.StatusOK, "Hello jiahaohong\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *altria.Context) {
		names := []string{"jiahaohong"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
