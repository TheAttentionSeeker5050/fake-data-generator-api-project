// main route endpoints with gin
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainRoutes is a function that receives a pointer to a routing group and returns a pointer to a routing group
func EndpointRoutes(router *gin.Engine) *gin.RouterGroup {

	// make routing group /
	endpointRouter := router.Group("/endpoints")
	{
		// make index route on /
		endpointRouter.GET("", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/index.html", gin.H{
				"title": "Hello World",
			})
		})

		// make a route to edit an endpoint
		endpointRouter.GET("/:id/edit", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/add-edit.html", gin.H{
				"title": "Hello World",
			})
		})

		// make a route to create an endpoint
		endpointRouter.GET("/create", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/add-edit.html", gin.H{
				"title": "Hello World",
			})
		})

		// make a route to delete an endpoint
		endpointRouter.GET("/:id/delete", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/delete.html", gin.H{
				"title": "Hello World",
			})
		})

		// make a route to view an endpoint
		endpointRouter.GET("/:id", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/details.html", gin.H{
				"title": "Hello World",
			})
		})

	}

	return endpointRouter
}
