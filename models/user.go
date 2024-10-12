package models

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// User model definition
type User struct {
	ID            string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName     string    `gorm:"size:50;not null"`
	LastName      string    `gorm:"size:50;not null"`
	Email         string    `gorm:"size:100;unique;not null"`
	HashedPassword string   `gorm:"size:200;not null"`
	IsActive      bool      `gorm:"default:true;not null"`
	IsSuperuser   bool      `gorm:"default:false;not null"`
	IsVerified    bool      `gorm:"default:false;not null"`
	CreatedAt     time.Time `gorm:"default:now()"`
	UpdatedAt     time.Time `gorm:"default:now()"`
}

// Claims model definition
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
