package controller

import (
	"net/http"

	"github.com/Tabed23/UserPostTracker/app/service"
	"github.com/Tabed23/UserPostTracker/app/types"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (c *PostController) GetPosts(ctx *gin.Context) {
	posts, err := c.postService.GetPosts(ctx, 10, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (c *PostController) GetPostByID(ctx *gin.Context) {
	postID := ctx.Param("id")
	post, err := c.postService.GetPostByID(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get post by ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})

}

func (c *PostController) CreatePost(ctx *gin.Context) {

	userID := ctx.Param("id")
	post := types.Post{}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post body"})
	}

	postId, err := c.postService.CreatePost(ctx, &post, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post_id": postId})
}

func (c *PostController) LikePost(ctx *gin.Context) {
	like := types.Like{}

	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid  ID"})
		return
	}

	err := c.postService.AddLike(ctx, like.PostID.Hex(), like.UserID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post like successfully"})
}

func (c *PostController) UnlikePost(ctx *gin.Context) {

	like := types.Like{}

	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid  ID"})
		return
	}

	err := c.postService.RemoveLike(ctx, like.PostID.Hex(), like.UserID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post unlike successfully"})
}

func (c *PostController) AddComment(ctx *gin.Context) {
	postID := ctx.Param("id")

	cmt := types.Comment{}

	if err := ctx.ShouldBindJSON(&cmt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := c.postService.AddComment(ctx, postID, cmt.UserID.Hex(), &cmt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"comment": res})
}

func (c *PostController) UpdatePost(ctx *gin.Context) {
	post := types.Post{}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post body"})
	}
	res, err := c.postService.UpdatePost(ctx, &post)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": res})
}
