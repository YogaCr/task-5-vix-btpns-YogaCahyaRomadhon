package middlewares

import (
	"net/http"
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/helpers"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := helpers.TokenValidation(request)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Unauthorized"))
			return
		}
		next(writer, request)
	}
}
