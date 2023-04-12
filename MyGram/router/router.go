package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "mygram/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		//CreateSocialMedia
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
