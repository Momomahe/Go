package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	socialmediaRouter := r.Group("socialmedia")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.GET("/:socialmediaID", controllers.GetOne)
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		socialmediaRouter.PUT("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/:photoID", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/:commentID", controllers.GetOneComment)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
