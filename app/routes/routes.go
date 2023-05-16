package routes

import (
	"github.com/Tabed23/UserPostTracker/app/controller"
	"github.com/Tabed23/UserPostTracker/app/database"
	"github.com/Tabed23/UserPostTracker/app/repo"
	"github.com/Tabed23/UserPostTracker/app/service"
)

func UrlMaps() {

	//userService := service.UserService()

	postRepo := repo.NewPostRepo(database.OpenCollection(database.Client, "postTask"))
	usrRepo := repo.NewUserRepo(database.OpenCollection(database.Client, "userTask"))

	postService := service.NewPostService(*postRepo)
	usrService := service.NewUserService(*usrRepo)

	postApi := controller.NewPostController(postService)
	usrApi := controller.NewUserController(usrService)

	v1 := r.Group("/v1")
	{
		v1.POST("/api/users", usrApi.CreateUser)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/api/posts", postApi.GetPosts)
		v2.GET("/api/posts/:id", postApi.GetPostByID)
		v2.POST("/api/posts", postApi.CreatePost)
		v2.PUT("/api/post", postApi.UpdatePost)
		v2.PUT("/api/posts/like", postApi.LikePost)
		v2.PUT("/api/posts/unlike", postApi.UnlikePost)
		v2.POST("/api/posts/:id/comments", postApi.AddComment)

	}

}
