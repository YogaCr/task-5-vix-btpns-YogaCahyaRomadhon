package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"task-5-vix-btpns-Yoga_Cahya_Romadhon/models"
	"time"
)

func GenerateJwt(user models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnv("JWT_SECRET")))
}

func TokenValidation(r *http.Request) error {
	tokenString := getTokenHeader(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
	}
	return nil
}

func getTokenHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	bearerString := "bearer "
	return bearerToken[len(bearerString):]
}
