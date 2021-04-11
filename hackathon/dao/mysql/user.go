package mysql

import (
	"gorm.io/gorm"
	"hackathon/models"
)

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
