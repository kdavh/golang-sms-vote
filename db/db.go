package db

import (
	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
)

func Init() (db *sqlx.DB) {
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

	d, err := sqlx.Connect("postgres", openCmd)
	if err != nil {
		log.Fatal(err)
	}

	return d
}
