package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name			string	`gorm:"not null;uniqueIndex" json:"name" form:"name" valid:"required~Social media name is required"`
	SocialMediaUrl 	string `gorm:"not null" json:"social_media_url" form:"social_media__url" valid:"required~Url of your social media_ is required"`
	UserID   		uint
	User     		*User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}

func (sm *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
