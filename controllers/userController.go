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

func UserGet(c *gin.Context) {
	db 			:= database.GetDB()
	userId, _ 	:= strconv.Atoi(c.Param("id"))
	Users 		:= models.User{}
	err 		:= db.Find(&Users, uint(userId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Users)
}

func UserGetList(c *gin.Context) {
	db := database.GetDB()
	Users := []models.User{}
	err := db.Find(&Users).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Users)
}

func UserCreate(c *gin.Context) {
	db := database.GetDB()

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	
	// Input Create User
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

func UserUpdate(c *gin.Context) {
	db := database.GetDB()

	var users models.User
	err := db.First(&users, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	// Input Update User
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()

	userId, _ := strconv.Atoi(c.Param("id"))
	Users := models.User{}

	err := db.Delete(Users, uint(userId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been successfully deleted",
	})
}
