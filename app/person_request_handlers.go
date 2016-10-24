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

func HandleCreate(c *gin.Context) {
	var db *gorm.DB
	db = c.MustGet("DB").(*gorm.DB)
	repo := PersonRepo{DB: db}

	var p Person
	c.BindJSON(&p)

	if errs := repo.Create(&p); errs != nil {
		errStrs := make([]string, len(errs))

		for i, err := range errs {
			errStrs[i] = err.Error()
		}

		c.JSON(422, map[string][]string{"errors": errStrs})
	} else {
		c.JSON(200, p)
	}
}
