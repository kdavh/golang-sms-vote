package app

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	FirstName string `db:"first_name" binding:"required"`
	LastName  string `db:"last_name" binding:"required"`
	Email     string `binding:"required"`
}
