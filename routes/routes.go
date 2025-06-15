package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// server.GET("/route-name", routeHandler)
	// server.GET("route-name/:id", routeHandler)

	// to enforce authentication on specific routes:

	// authenticated := server.Group("/")
	// authenticated.Use(middlewares.Authenticate)

	// authenticated.POST("/route-name", routeHandler)
	// authenticated.PUT("/route-name/:id", routeHandler)
	fmt.Print("Do something")
}
