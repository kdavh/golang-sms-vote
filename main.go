package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	env := "development"
	cfgFile := filepath.Join("db", "dbconf.yml")

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	openCmd, _ := f.Get(fmt.Sprintf("%s.open", env))

	db, err := sqlx.Connect("postgres", openCmd)

	if err != nil {
		log.Fatalln(err)
	}

	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)

	jason = Person{}
	err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", jason)

	places := []Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}

	// Loop through rows using only one struct
	place := Place{}
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}

	// routing
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
