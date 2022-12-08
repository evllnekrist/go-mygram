package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	GormModel
	Name			string	`gorm:"not null;uniqueIndex" json:"name" form:"name" valid:"required~Social media name is required"`
	SocialMediaUrl 	string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Url of your social media_ is required"`
	UserID   		uint `json:"user_id"`
	User     		*User
}

func (sm *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}

func (sm *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
