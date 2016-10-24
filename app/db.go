package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func InitDB() (db *gorm.DB) {
	openCmd := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable")

	d, err := gorm.Open("postgres", openCmd)
	if err != nil {
		log.Fatal(err)
	}

	// config
	d.SingularTable(true)
	d.AutoMigrate(&Person{})

	return d
}

func DBInjectionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
