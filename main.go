package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	// prepare jsoniter response
	router.GET("/api", func(c *gin.Context) {

		var data = struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "John",
			Age:  30,
		}

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Marshal(&data)
		c.JSON(200, data)
	})

	// make hello world h1 tag on / in plain text
	router.GET("/", func(c *gin.Context) {

		// render h1 tag
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
