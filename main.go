package main

import (
	"fmt"
	"geb/lib/geb"
)

func main() {

	r := geb.New()

	r.GET("/", func(c *geb.Context) {
		c.String(200, "hello")
	})

	r.GET("/query", func(c *geb.Context) {
		c.String(200, c.Query("code"))
	})

	r.POST("/form", func(c *geb.Context) {
		c.String(200, c.PostForm("code"))
	})

	fmt.Println("server start")
	r.Run(":9999")

}
