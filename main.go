package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kdavh/gin-sms-vote/db"
	"github.com/kdavh/gin-sms-vote/person"
)

func main() {
	d := db.Init()
	fmt.Println(d)

	// initialize the resource and inject our db connection
	repo := &person.Repo{DB: d}

	// routing
	r := gin.Default()

	r.POST("/person", repo.Create)
	r.GET("/person/:id", repo.Get)
	r.DELETE("/person/:id", repo.Delete)

	r.Run() // listen and server on 0.0.0.0:8080
}
