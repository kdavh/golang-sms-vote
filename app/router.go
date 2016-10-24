package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(d *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(DBInjectionMiddleware(d))

	r.GET("/person/:id", HandleShow)
	r.POST("/person", HandleCreate)

	return r
}
