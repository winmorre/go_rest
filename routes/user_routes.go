package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	user "rest/controllers/user"
)

func StartGin() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.GET("/user/:id", user.GetUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8080")
}
