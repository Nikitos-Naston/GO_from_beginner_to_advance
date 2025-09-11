package routers

import (
	handlers "Ginogorm/Handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("article", handlers.GetallArticles)
	apiV1Group.POST("article", handlers.PostNewArticl)
	apiV1Group.GET("article/:id", handlers.GetArticlebyID)
	apiV1Group.PUT("article/:id", handlers.UpdateArcticleById)
	apiV1Group.DELETE("article/:id", handlers.DeleteArctilebyId)

	//apiUserGroup := router.Group("/user")

	return router
}
