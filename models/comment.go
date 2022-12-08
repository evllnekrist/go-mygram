package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message  	string `gorm:"not null" json:"message" form:"message" valid:"required~Message of your comment is required"`
	PhotoID   	uint `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo ID of your comment is required"`
	Photo     	*Photo
	UserID   	uint `json:"user_id"`
	User     	*User
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
