package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserIDAndRoles struct {
	UserId string   `json:"userId"`
	Roles  []string `json:"roles"`
}
type MyCustomClaims struct {
	UserIDAndRoles
	jwt.StandardClaims
}

type UserDataString string

var UserDataKey UserDataString = "userData"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		fmt.Println("Logging the request", time.Now())
		token := c.Request.Header.Get("Authorization")
		userIdAndRoles := VerifyToken(token)

		ctx = context.WithValue(c.Request.Context(), UserDataKey, userIdAndRoles)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

var mySigningKey = []byte("AllYourBase")

func SignToken(userId string) *string {
	claims := MyCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour) * 3000,
			Issuer:    "api",
		},
	}
	claims.Roles = []string{"ADMIN"}
	claims.UserId = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, tokenErr := token.SignedString(mySigningKey)
	if tokenErr != nil {
		fmt.Println(tokenErr)
		return nil
	}
	return &ss
}

func VerifyToken(tokenString string) *UserIDAndRoles {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return &UserIDAndRoles{
			Roles:  claims.Roles,
			UserId: claims.UserId,
		}

	} else {
		fmt.Println(err)
		return nil
	}
}
