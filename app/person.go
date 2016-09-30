package app

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// CreatedAt int32  `db:"created_at"`
	Email string
}

//func (Person) TableName() string {
//return "person"
//}
