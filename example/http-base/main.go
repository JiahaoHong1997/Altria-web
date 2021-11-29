package main

import (
	"fmt"
	"github.com/JiahaoHong1997/altria-web"
	"net/http"
)

func main() {
	r := altria_web.Saber()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf()
	})
}
