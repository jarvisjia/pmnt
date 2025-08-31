package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jarvisjia/pmnt/go/task4/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var JwtSecret = []byte("your-secret-key")

type UserService struct {
	BaseService
}

func (u *UserService) RegisterPage(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("register page")
	ctx.HTML(http.StatusOK, "users/register.tmpl", nil)
}

func (u *UserService) Register(ctx *gin.Context, db *gorm.DB) {
	username := ctx.PostForm("username")
	if username == "" {
		u.error(ctx, http.StatusBadRequest, "username is required")
		return
	}
	password := ctx.PostForm("password")
	if password == "" {
		u.error(ctx, http.StatusBadRequest, "password is required")
		return
	}
	fmt.Println("register form:", username, password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		u.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	pwd := string(hashedPassword)
	user := model.User{
		Username: username,
		Password: pwd,
		Email:    fmt.Sprintf("%s@xxx.com", username),
	}
	fmt.Println("user:", user)
	if err := db.Debug().Create(&user).Error; err != nil {
		u.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	u.success(ctx, gin.H{"message": "registered successfully"})
}

func (u *UserService) LoginPage(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("login page")
	ctx.HTML(http.StatusOK, "users/login.tmpl", nil)
}

func (u *UserService) Login(ctx *gin.Context, db *gorm.DB) {
	username := ctx.PostForm("username")
	if username == "" {
		u.error(ctx, http.StatusBadRequest, "username is required")
		return
	}
	password := ctx.PostForm("password")
	if password == "" {
		u.error(ctx, http.StatusBadRequest, "password is required")
		return
	}
	fmt.Println("login form:", username, password)

	user := model.User{}
	if err := db.Debug().Where("username = ?", username).First(&user).Error; err != nil {
		u.error(ctx, http.StatusUnauthorized, "invalid username")
		return
	}
	fmt.Println("user from db:", user)

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		u.error(ctx, http.StatusUnauthorized, "invalid password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   user.ID,
		"username": user.Username,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		u.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	u.success(ctx, gin.H{"message": "login success", "token": tokenString})
}
