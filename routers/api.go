package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/controllers"
	"go-rest-api/middlewares"
)

func StartServer() *gin.Engine{
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", controllers.UserGetList)
		userRouter.GET("/:id", controllers.UserGet)
		userRouter.POST("/register", controllers.UserCreate)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:id", middlewares.UserAuthorization(), controllers.UserUpdate)
		userRouter.DELETE("/:id", middlewares.UserAuthorization(), controllers.UserDelete)
	}
	
	photoRouter := r.Group("/photos")
	{
		photoRouter.GET("/", controllers.PhotoGetList)
		photoRouter.GET("/:id", controllers.PhotoGet)
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", middlewares.PhotoAuthorization(), controllers.PhotoCreate)
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.GET("/", controllers.CommentGetList)
		commentRouter.GET("/:id", controllers.CommentGet)
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", middlewares.CommentAuthorization(), controllers.CommentCreate)
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.CommentUpdate)
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.CommentDelete)
	}

	socialMediaRouter := r.Group("/social-medias")
	{
		socialMediaRouter.GET("/", controllers.SocialMediaGetList)
		socialMediaRouter.GET("/:id", controllers.SocialMediaGet)
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", middlewares.SocialMediaAuthorization(), controllers.SocialMediaCreate)
		socialMediaRouter.PUT("/:id", middlewares.SocialMediaAuthorization(), controllers.SocialMediaUpdate)
		socialMediaRouter.DELETE("/:id", middlewares.SocialMediaAuthorization(), controllers.SocialMediaDelete)
	}

	return r
}