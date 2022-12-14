package middlewares

import (
	// "fmt"
	"net/http"
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-rest-api/database"
	"go-rest-api/models"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authName := "User"
		db := database.GetDB()
		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		// User := models.User{}
		
		if c.Request.Method != "POST" {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			Users := models.User{}
			err = db.Select("id").First(&Users, uint(id)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

			if userID == Users.ID {
				return
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data. "+authName+" belongs to somebody else",
				})
				return
			}
		}else if c.Request.Method == "POST" {
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authName := "Photo"
		db := database.GetDB()
		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		// User := models.User{}
		
		if c.Request.Method != "POST" {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			Photos := models.Photo{}
			err = db.Select("user_id").First(&Photos, uint(id)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

			if userID == Photos.UserID {
				return
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data. "+authName+" belongs to somebody else",
				})
				return
			}
		}else if c.Request.Method == "POST" {
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authName := "Comment"
		db := database.GetDB()
		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		// User := models.User{}
		
		if c.Request.Method != "POST" {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			Comments := models.Comment{}
			err = db.Select("user_id").First(&Comments, uint(id)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

			if userID == Comments.UserID {
				return
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data. "+authName+" belongs to somebody else",
				})
				return
			}
		}else if c.Request.Method == "POST" {
			return
		}

		c.Next()
	}
}

func SocialmediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authName := "Social media"
		db := database.GetDB()
		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		// User := models.User{}
		
		if c.Request.Method != "POST" {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": "invalid parameter",
				})
				return
			}

			Socialmedias := models.Socialmedia{}
			err = db.Select("user_id").First(&Socialmedias, uint(id)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doesn't exist",
				})
				return
			}

			if userID == Socialmedias.UserID {
				return
			}else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to access this data. "+authName+" belongs to somebody else",
				})
				return
			}
		}else if c.Request.Method == "POST" {
			return
		}

		c.Next()
	}
}