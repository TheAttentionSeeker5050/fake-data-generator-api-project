// This package contains user REST API routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiUserRoutes is a group of routes for user REST API requests
func ApiUserRoutes(router *gin.Engine) *gin.RouterGroup {

	// routing group /api/user
	apiUserRouter := router.Group("/api/user")
	{

		// login route POST
		apiUserRouter.POST("/login", func(c *gin.Context) {
			// render page
			c.JSON(http.StatusOK, gin.H{
				"message": "User is logged in",
			})
		})

		// logout route POST
		apiUserRouter.POST("/logout", func(c *gin.Context) {
			// render page
			c.JSON(http.StatusOK, gin.H{
				"message": "User is logged out",
			})
		})

		// register route POST
		apiUserRouter.POST("/register", func(c *gin.Context) {
			// render page
			c.JSON(http.StatusOK, gin.H{
				"message": "User is registered",
			})
		})

		// change profile route PUT
		apiUserRouter.PUT("/profile", func(c *gin.Context) {
			// render page
			c.JSON(http.StatusOK, gin.H{
				"message": "Profile is changed",
			})
		})
	}

	return apiUserRouter
}
