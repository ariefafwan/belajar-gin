package routers

import (
	"belajar-gin/config"
	"belajar-gin/internal/controllers"
	"belajar-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func registerBioskopRoutes(router *gin.Engine) {
	repo := repositories.NewBioskopRepository(config.DB)
	controller := controllers.NewBioskopController(repo)

	bioskops := router.Group("/bioskops")
	{
		bioskops.GET("", controller.GetAll)
		bioskops.GET("/:id", controller.GetByID)
		bioskops.POST("", controller.Create)
		bioskops.PUT("/:id", controller.Update)
		bioskops.DELETE("/:id", controller.Delete)
	}
}
