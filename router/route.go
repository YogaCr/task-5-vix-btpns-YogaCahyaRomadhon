package router

import (
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/controllers"
)

func InitRoutes(s *controllers.Server) {
	s.Router.POST("register", s.CreateUser)
	s.Router.GET("login", s.Login)
}
