package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jarvisjia/pmnt/go/task4/service"
	"net/http"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("JWT middleware")

		// Get the token from the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return service.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			ctx.Abort()
			return
		}
		fmt.Println("Token is valid:", token)
		fmt.Println("Claims:", token.Claims)
		ctx.Set("userid", claims["userid"])
		ctx.Set("username", claims["username"])

		// Token is valid, proceed to the next handler
		ctx.Next()
	}
}
