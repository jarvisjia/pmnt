package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jarvisjia/pmnt/go/task4/model"
	"gorm.io/gorm"
)

type PostService struct {
	BaseService
}

func (p *PostService) ListPost(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("list post")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)

	posts := []model.Post{}
	if err := db.Where("user_id=?", userid).Find(&posts).Error; err != nil {
		p.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	// ctx.HTML(http.StatusOK, "posts/list.tmpl", gin.H{"posts": posts})
	p.success(ctx, posts)
}

func (p *PostService) DetailPost(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("detail post")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	id := ctx.Param("id")
	fmt.Println("post id:", id)

	post := model.Post{}
	if err := db.Where("id=? and user_id=?", id, userid).First(&post).Error; err != nil {
		p.error(ctx, http.StatusBadRequest, "post not found")
		return
	}
	// ctx.HTML(http.StatusOK, "posts/detail.tmpl", gin.H{"post": post})
	p.success(ctx, post)
}

func (p *PostService) CreatePost(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("create post")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	fmt.Println("title:", title, "content:", content)
	if title == "" || content == "" {
		p.error(ctx, http.StatusBadRequest, "title and content are required")
		return
	}

	post := model.Post{
		Title:   title,
		Content: content,
		UserID:  uint(userid.(float64)),
	}
	if err := db.Create(&post).Error; err != nil {
		p.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	p.success(ctx, gin.H{"message": "post created successfully"})
}

func (p *PostService) EditPost(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("edit post")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	id := ctx.PostForm("id")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	if id == "" || title == "" || content == "" {
		p.error(ctx, http.StatusBadRequest, "id, title and content are required")
		return
	}
	post := model.Post{}
	if err := db.Where("id = ? and user_id=?", id, userid).First(&post).Error; err != nil {
		p.error(ctx, http.StatusBadRequest, "post not found")
		return
	}
	post.Title = title
	post.Content = content
	if err := db.Save(&post).Error; err != nil {
		p.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	p.success(ctx, gin.H{"message": "post updated successfully"})
}

func (p *PostService) DeletePost(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("delete post")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	id := ctx.Param("id")
	fmt.Println("post id:", id)
	if id == "" {
		p.error(ctx, http.StatusBadRequest, "id is required")
		return
	}
	post := model.Post{}
	if err := db.Where("id = ? and user_id=?", id, userid).First(&post).Error; err != nil {
		p.error(ctx, http.StatusBadRequest, "post not found")
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		p.error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	p.success(ctx, gin.H{"message": "post deleted successfully"})
}
