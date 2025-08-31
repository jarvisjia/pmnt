package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jarvisjia/pmnt/go/task4/middleware"
	"github.com/jarvisjia/pmnt/go/task4/service"
	"gorm.io/gorm"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {
	user := r.Group("/user")
	{
		service := &service.UserService{}
		user.GET("/rp", func(ctx *gin.Context) { service.RegisterPage(ctx, db) })
		user.POST("/register", func(ctx *gin.Context) { service.Register(ctx, db) })
		user.GET("/lp", func(ctx *gin.Context) { service.LoginPage(ctx, db) })
		user.POST("/login", func(ctx *gin.Context) { service.Login(ctx, db) })
	}
}

func PostRouter(r *gin.Engine, db *gorm.DB) {
	post := r.Group("/post")
	post.Use(middleware.JwtMiddleware()).Use(middleware.LoginMiddleware(db))
	{
		service := &service.PostService{}
		post.GET("list", func(ctx *gin.Context) { service.ListPost(ctx, db) })
		post.GET("/detail/:id", func(ctx *gin.Context) { service.DetailPost(ctx, db) })
		post.POST("/create", func(ctx *gin.Context) { service.CreatePost(ctx, db) })
		post.POST("/edit", func(ctx *gin.Context) { service.EditPost(ctx, db) })
		post.POST("/delete/:id", func(ctx *gin.Context) { service.DeletePost(ctx, db) })
	}
}

func CommentRouter(r *gin.Engine, db *gorm.DB) {
	comment := r.Group("/comment")
	comment.Use(middleware.JwtMiddleware()).Use(middleware.LoginMiddleware(db))
	{
		service := &service.CommentService{}
		comment.POST("/create", func(ctx *gin.Context) { service.CreateComment(ctx, db) })
		comment.GET("/list/:postid", func(ctx *gin.Context) { service.ListComment(ctx, db) })
	}
}
