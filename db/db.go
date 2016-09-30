package db

import (
	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"path"
	"path/filepath"
	"runtime"
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
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	cfgFile := filepath.Join(path.Dir(filename), "dbconf.yml")

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	openCmd, _ := f.Get(fmt.Sprintf("%s.open", dbEnv))

	d, err := gorm.Open("postgres", openCmd)
	if err != nil {
		log.Print("hello")
		log.Fatal(err)
	}

	// config
	d.SingularTable(true)

	return d
}

func InjectionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
