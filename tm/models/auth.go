package models

import (
	"time"
	"tm/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique" json:"username" validate:"min=6,max=50" binding:"required"`
	Password  string    `json:"password" validate:"min=6,max=50" binding:"required"`
	Active    int       `gorm:"active"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	DB.Where("username = ? AND active = ?", username, 1).First(&auth)
	if auth.ID > 0 && checkPassword(password, auth.Password) {
		return true
	}
	return false
}

func GenPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logging.Error(err.Error())
	}
	return string(hash)
}

func checkPassword(password string, source string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(source), []byte(password)); err == nil {
		return true
	}
	return false
}
