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

func PhotoGet(c *gin.Context) {
	db 			:= database.GetDB()
	photoId, _ 	:= strconv.Atoi(c.Param("id"))
	Photos 		:= models.Photo{}
	err 		:= db.Find(&Photos, uint(photoId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photos)
}

func PhotoGetList(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	err := db.Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photos)
}

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	input.UserID = userID
	
	// Input Create Photo
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

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	var photos models.Photo
	err := db.First(&photos, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	Result := map[string]interface{}{}
	SqlStatement := "Update photos SET title = ?, caption = ?, photo_url = ?, updated_at = ?  WHERE id = ? RETURNING id, title, caption, photo_url, created_at, user_id"
	err2 := db.Raw(
		SqlStatement,
		input.Title, input.Caption, input.PhotoUrl, time.Now(), c.Param("id"),
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

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()

	photoId, _ := strconv.Atoi(c.Param("id"))
	Photos := models.Photo{}

	err := db.Delete(Photos, uint(photoId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo has been successfully deleted",
	})
}
