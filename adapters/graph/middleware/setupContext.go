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

type AuthorizationString string

var UserDataKey UserDataString = "userData"
var AuthorizationHeader AuthorizationString = "Authorization"

func SetupContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		token := c.Request.Header.Get("Authorization")
		ctx = context.WithValue(c.Request.Context(), AuthorizationHeader, token)
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
	if tokenString == "" {
		return nil
	}
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return &UserIDAndRoles{
			Roles:  claims.Roles,
			UserId: claims.UserId,
		}

	} else {
		fmt.Println("Unable to verify token")
		return nil
	}
}
