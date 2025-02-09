package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint
	User    *User
	PhotoID uint `json:"photo_id" form:"photo_id" validation:"required"`
	Photo   *Photo
	Message string `json:"message" form:"message" validation:"required"`
}

func (p *Comment) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate() (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	err = nil
	return
}
