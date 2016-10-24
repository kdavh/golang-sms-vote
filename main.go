package main

import (
	"github.com/kdavh/gin-sms-vote/app"
)

func main() {
	d := app.InitDB()
	r := app.InitRouter(d)

	r.Run() // listen and server on 0.0.0.0:8080
}
