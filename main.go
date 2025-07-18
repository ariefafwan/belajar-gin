package main

import (
	"belajar-gin/config"
	"belajar-gin/routers"
	"fmt"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.InitValidator()

	var router = routers.SetupRouter()
	router.Run(":8080")
	fmt.Println("Server berjalan di")
}