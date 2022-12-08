package controllers

import (
	"net/http"
	"strconv"
	"time"
	// "fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-rest-api/database"
	"go-rest-api/models"
)

func SocialmediaGet(c *gin.Context) {
	db 			:= database.GetDB()
	socialmediaId, _ 	:= strconv.Atoi(c.Param("id"))
	Socialmedias 		:= models.Socialmedia{}
	err 		:= db.Find(&Socialmedias, uint(socialmediaId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Socialmedias)
}

func SocialmediaGetList(c *gin.Context) {
	db := database.GetDB()
	Socialmedias := []models.Socialmedia{}
	err := db.Find(&Socialmedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Socialmedias)
}

func SocialmediaCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var input models.Socialmedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	input.UserID = userID
	
	// Input Create Socialmedia
	err := db.Debug().Create(&input).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

func SocialmediaUpdate(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	var socialmedias models.Socialmedia
	err := db.First(&socialmedias, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Socialmedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	
	Result := map[string]interface{}{}
	SqlStatement := "Update socialmedia SET name = ?, social_media_url = ?, updated_at = ?  WHERE id = ? RETURNING id, name, social_media_url, created_at, user_id"
	err2 := db.Raw(
		SqlStatement,
		input.Name, input.SocialMediaUrl, time.Now(), c.Param("id"),
	).Scan(&Result).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Result,
		"code" : http.StatusOK,
	})
}

func SocialmediaDelete(c *gin.Context) {
	db := database.GetDB()

	socialmediaId, _ := strconv.Atoi(c.Param("id"))
	Socialmedias := models.Socialmedia{}

	err := db.Delete(Socialmedias, uint(socialmediaId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media has been successfully deleted",
	})
}
