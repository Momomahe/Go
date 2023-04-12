package router

import (
	_ "golang-crud-gin/docs"
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

// @title User API
// @version 1.0
// @description This a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licences/LICENCE-2.0.html
// @host localhost:3000
// @BasePath /
func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("user")
	{
		//Register
		userRouter.POST("/register", controllers.Register)
		//Login
		userRouter.POST("/login", controllers.Login)
	}

	socialmediaRouter := r.Group("socialmedia")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.GET("/:socialmediaID", controllers.GetOne)
		socialmediaRouter.GET("/", controllers.GetAllComment)
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		socialmediaRouter.PUT("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialmediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/:photoID", controllers.GetOnePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/:commentID", controllers.GetOneComment)
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
