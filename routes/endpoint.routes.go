// main route endpoints with gin
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EndpointRoutes is a routing group for endpoints website pages, these are the pages that the user will interact with
// the usr will be able to create, edit, delete and view endpoints, these will then be used to make requests
// from their frontend applications easily, point and click, no code on backend required
func EndpointRoutes(router *gin.Engine) *gin.RouterGroup {

	// make routing group /
	endpointRouter := router.Group("/endpoints")
	{
		// page to list all endpoints
		endpointRouter.GET("", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/index.html", gin.H{
				"title": "Hello World",
			})
		})

		// page to edit an endpoint
		endpointRouter.GET("/:id/edit", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/add-edit.html", gin.H{
				"title": "Hello World",
			})
		})

		// page to create an endpoint
		endpointRouter.GET("/create", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/add-edit.html", gin.H{
				"title": "Hello World",
			})
		})

		// page to delete an endpoint
		endpointRouter.GET("/:id/delete", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/delete.html", gin.H{
				"title": "Hello World",
			})
		})

		// page to view an endpoint
		endpointRouter.GET("/:id", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "endpoints/details.html", gin.H{
				"title": "Hello World",
			})
		})

	}

	return endpointRouter
}
