// This package contains pages routes for the website, these will return the pages to the browser
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainRoutes is a routing group for main routes, these consist of the index, authentication and profile routes
// profile and logout routes are protected and require authentication and ownership of the profile
func MainRoutes(router *gin.Engine) *gin.RouterGroup {

	// make routing group /
	mainRouter := router.Group("/")
	{
		// make index route on /
		mainRouter.GET("", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "main/index.html", gin.H{
				"title": "Hello World",
			})
		})

		// make login route on /login
		mainRouter.GET("/login", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "main/login.html", gin.H{
				// "title": "Hello World",
			})
		})

		// make register route on /register
		mainRouter.GET("/register", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "main/register.html", gin.H{
				// "title": "Hello World",
			})
		})

		// make logout route on /logout
		mainRouter.GET("/logout", func(c *gin.Context) {
			// render page
			c.HTML(http.StatusOK, "main/logout.html", gin.H{
				// "title": "Hello World",
			})
		})

		// make profile route on /profile
		mainRouter.GET("/profile", func(c *gin.Context) {
			c.HTML(http.StatusOK, "main/profile.html", gin.H{
				// "title": "Hello World",
			})
		})
	}

	return mainRouter
}
