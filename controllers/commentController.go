package controllers

import (
	"net/http"
	"strconv"
	// "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	
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

	// Input Update Comment
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
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
