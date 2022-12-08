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
		commentRouter.POST("/", controllers.CommentCreate)
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.CommentUpdate)
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.CommentDelete)
	}

	socialmediaRouter := r.Group("/social-medias")
	{
		socialmediaRouter.GET("/", controllers.SocialmediaGetList)
		socialmediaRouter.GET("/:id", controllers.SocialmediaGet)
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.POST("/", middlewares.SocialmediaAuthorization(), controllers.SocialmediaCreate)
		socialmediaRouter.PUT("/:id", middlewares.SocialmediaAuthorization(), controllers.SocialmediaUpdate)
		socialmediaRouter.DELETE("/:id", middlewares.SocialmediaAuthorization(), controllers.SocialmediaDelete)
	}

	return r
}