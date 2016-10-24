package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kdavh/gin-sms-vote/app"
	"github.com/kdavh/gin-sms-vote/db"
)

func main() {
	d := db.Init()
	fmt.Println(d)

	// routing
	r := gin.Default()
	r.Use(db.InjectionMiddleware(d))

	r.GET("/person/:id", app.HandleShow)
	r.GET("/hello", handleHello)
	//r.POST("/person", repo.Create)
	//r.DELETE("/person/:id", repo.Delete)

	r.Run() // listen and server on 0.0.0.0:8080
}

func handleHello(c *gin.Context) {
    c.JSON(200, "{\"hello\": \"person\"}")
}
