package router

import (
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/controllers"
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/middlewares"
)

func InitRoutes(s *controllers.Server) {
	s.Router.POST("register", s.CreateUser)
	s.Router.GET("login", s.Login)

	needAuth := s.Router.Group("/")
	needAuth.Use(middlewares.AuthMiddleware(s.DB))
	{
		needAuth.PUT("users/:userId", s.UpdateUser)
		needAuth.DELETE("users/:userId", s.DeleteUser)
	}
}
