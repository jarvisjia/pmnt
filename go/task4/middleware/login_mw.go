package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarvisjia/pmnt/go/task4/model"
	"gorm.io/gorm"
)

func LoginMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("login middleware")

		username, ok := ctx.Get("username")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username not found in token"})
			ctx.Abort()
			return
		}
		userid, ok := ctx.Get("userid")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "userid not found in token"})
			ctx.Abort()
			return
		}
		fmt.Println("username from token:", username)
		fmt.Println("userid from token:", userid)

		user := model.User{}
		if err := db.Where("id = ? and username = ?", userid, username).First(&user).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			ctx.Abort()
			return
		}
		fmt.Println("user from db:", user)

		// ctx.Set("user", user)
		// User is authenticated, proceed to the next handler
		ctx.Next()
	}
}
