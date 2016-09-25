package db

import (
	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"path/filepath"
)

func Init() (db *gorm.DB) {
	var dbEnv string
	if gin.IsDebugging() {
		dbEnv = "development"
	} else {
		dbEnv = "production"
	}

	cfgFile := filepath.Join("db", "dbconf.yml")

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	openCmd, _ := f.Get(fmt.Sprintf("%s.open", dbEnv))

	d, err := gorm.Open("postgres", openCmd)
	if err != nil {
		log.Fatal(err)
	}

	// config
	d.SingularTable(true)

	return d
}
