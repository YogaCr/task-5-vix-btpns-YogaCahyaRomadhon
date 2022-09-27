package main

import (
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/controllers"
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/router"
)

func main() {
	server := controllers.Server{}
	server.Initialize()
	router.InitRoutes(&server)
	server.Run(8000)
}
