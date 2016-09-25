package person

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repo struct {
	DB *sqlx.DB
}

func (repo *Repo) Get(c *gin.Context) {
	var m Model

	_ = repo.DB.Get(&m, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", m)

	err := repo.DB.Get(&m, "SELECT * FROM person WHERE id = $1", c.Param("id"))
	log.Print(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, m)
	}

	log.Print(m)
}

func (repo *Repo) Create(c *gin.Context) {
	var m Model

	c.Bind(&m)

	m.FirstName = "abc"
	//m.Created = int32(time.Now().Unix())

	//repo.db.Save(&m)
	c.JSON(201, m)
}

func (repo *Repo) Delete(c *gin.Context) {
}

//func (tr *TodoResource) GetTodo(c *gin.Context) {
//idStr := c.Params.ByName("id")
//idInt, _ := strconv.Atoi(idStr)
//id := int32(idInt)

//var todo api.Todo

//if tr.db.First(&todo, id).RecordNotFound() {
//c.JSON(404, gin.H{"error": "not found"})
//} else {
//c.JSON(200, todo)
//}
//}

//func (tr *TodoResource) DeleteTodo(c *gin.Context) {
//idStr := c.Params.ByName("id")
//idInt, _ := strconv.Atoi(idStr)
//id := int32(idInt)

//var todo api.Todo

//if tr.db.First(&todo, id).RecordNotFound() {
//c.JSON(404, gin.H{"error": "not found"})
//} else {
//tr.db.Delete(&todo)
//c.Data(204, "application/json", make([]byte, 0))
//}
//}
//func Get(db *sqlx.DB) {
//people := []Person{}

//db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
//jason, john := people[0], people[1]

//fmt.Printf("%#v\n%#v", jason, john)

//jason = Person{}
//err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
//fmt.Printf("%#v\n", jason)

//// Loop through rows using only one struct
//place := Place{}
//rows, err := db.Queryx("SELECT * FROM place")
//for rows.Next() {
//err := rows.StructScan(&place)
//if err != nil {
//log.Fatalln(err)
//}
//fmt.Printf("%#v\n", place)
//}

//}
