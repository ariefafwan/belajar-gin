package routers

import (
	"belajar-gin/config"
	"belajar-gin/internal/controllers"
	"belajar-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	// api := router.Group("/api")
	// {
		bioskopRepo := repositories.NewBioskopRepository(config.DB)
		bioskopController := controllers.NewBioskopController(bioskopRepo)

		bioskop := router.Group("/bioskops")
		{
			bioskop.GET("/", bioskopController.GetAll)
			bioskop.GET("/:id", bioskopController.GetByID)
			bioskop.POST("/", bioskopController.Create)
			bioskop.PUT("/:id", bioskopController.Update)
			bioskop.DELETE("/:id", bioskopController.Delete)
		}
	// }
}
