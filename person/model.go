package person

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// CreatedAt int32  `db:"created_at"`
	Email string
}

func (Model) TableName() string {
	return "person"
}
