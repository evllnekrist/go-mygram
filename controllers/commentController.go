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

func CommentGet(c *gin.Context) {
	db 			:= database.GetDB()
	commentId, _ 	:= strconv.Atoi(c.Param("id"))
	Comments 		:= models.Comment{}
	err 		:= db.Find(&Comments, uint(commentId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)
}

func CommentGetList(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	err := db.Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)
}

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	input.UserID = userID
	
	// Input Create Comment
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

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	var comments models.Comment
	err := db.First(&comments, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	Result := map[string]interface{}{}
	SqlStatement := "Update comments SET message = ?, updated_at = ?  WHERE id = ? RETURNING id, message, photo_id, created_at, user_id"
	err2 := db.Raw(
		SqlStatement,
		input.Message, time.Now(), c.Param("id"),
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

func CommentDelete(c *gin.Context) {
	db := database.GetDB()

	commentId, _ := strconv.Atoi(c.Param("id"))
	Comments := models.Comment{}

	err := db.Delete(Comments, uint(commentId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment has been successfully deleted",
	})
}
