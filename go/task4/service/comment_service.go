package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarvisjia/pmnt/go/task4/model"
	"gorm.io/gorm"
	"net/http"
)

type CommentService struct {
	BaseService
}

func (c *CommentService) CreateComment(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("create comment")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	content := ctx.PostForm("content")
	postid := ctx.PostForm("postid")
	fmt.Println("content:", content, "postid:", postid)
	if content == "" || postid == "" {
		c.error(ctx, http.StatusBadRequest, "content and postid are required")
		return
	}

	post := model.Post{}
	if err := db.Where("id = ? and user_id=?", postid, userid).First(&post).Error; err != nil {
		c.error(ctx, http.StatusBadRequest, "post not found")
		return
	}
	comment := model.Comment{
		Content: content,
		UserID:  uint(userid.(float64)),
		PostID:  post.ID,
	}
	if err := db.Create(&comment).Error; err != nil {
		c.error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.success(ctx, gin.H{"message": "comment created successfully"})
}

func (c *CommentService) ListComment(ctx *gin.Context, db *gorm.DB) {
	fmt.Println("list comment")

	userid, _ := ctx.Get("userid")
	fmt.Println("userid from token:", userid)
	postid := ctx.Param("postid")
	fmt.Println("post id:", postid)

	post := model.Post{}
	if err := db.Where("id = ? and user_id=?", postid, userid).First(&post).Error; err != nil {
		c.error(ctx, http.StatusBadRequest, "post not found")
		return
	}
	var comments []model.Comment
	if err := db.Where("post_id = ?", postid).Find(&comments).Error; err != nil {
		c.error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.success(ctx, comments)
}
