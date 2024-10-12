package services

import (
    "gorm.io/gorm"
    "go-boilerplate/models"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
    db = database
}

func CreateUser(user *models.User) error {
    return db.Create(user).Error
}

func GetUserByEmail(email string) (models.User, error) {
    var user models.User
    err := db.Where("email = ?", email).First(&user).Error
    return user, err
}
