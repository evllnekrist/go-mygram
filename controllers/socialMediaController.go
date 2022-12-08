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

func SocialMediaGet(c *gin.Context) {
	db 			:= database.GetDB()
	socialMediaId, _ 	:= strconv.Atoi(c.Param("id"))
	SocialMedias 		:= models.SocialMedia{}
	err 		:= db.Find(&SocialMedias, uint(socialMediaId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedias)
}

func SocialMediaGetList(c *gin.Context) {
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}
	err := db.Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedias)
}

func SocialMediaCreate(c *gin.Context) {
	db := database.GetDB()

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	
	// Input Create SocialMedia
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

func SocialMediaUpdate(c *gin.Context) {
	db := database.GetDB()

	var socialMedias models.SocialMedia
	err := db.First(&socialMedias, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	// Input Update SocialMedia
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

func SocialMediaDelete(c *gin.Context) {
	db := database.GetDB()

	socialMediaId, _ := strconv.Atoi(c.Param("id"))
	SocialMedias := models.SocialMedia{}

	err := db.Delete(SocialMedias, uint(socialMediaId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SocialMedia has been successfully deleted",
	})
}
