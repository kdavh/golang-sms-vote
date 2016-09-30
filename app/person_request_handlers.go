package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func HandleShow(c *gin.Context) {
	var db *gorm.DB
	db = c.MustGet("DB").(*gorm.DB)
	repo := PersonRepo{DB: db}

	if m, err := repo.Get(c.Param("id")); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, m)
	}
}
