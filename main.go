package main

import (
	"belajar-gin/config"
	"belajar-gin/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()
	routers.SetupRouters(r)

	r.Run(":8080")
	fmt.Println("Server berjalan di")
}