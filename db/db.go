package db

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kdavh/gin-sms-vote/app"
	"log"
	"os"
)

func Init() (db *gorm.DB) {
	var dbEnv string
	if gin.IsDebugging() {
		dbEnv = "development"
	} else {
		dbEnv = "production"
	}

	return InitEnv(dbEnv)
}

func InitEnv(dbEnv string) (db *gorm.DB) {
	openCmd := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable")

	d, err := gorm.Open("postgres", openCmd)
	if err != nil {
		log.Fatal(err)
	}

	// config
	d.SingularTable(true)
	d.AutoMigrate(&app.Person{})

	return d
}

func InjectionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
