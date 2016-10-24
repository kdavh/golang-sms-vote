package app

import (
	"github.com/jinzhu/gorm"
	"log"
)

type PersonRepo struct {
	DB *gorm.DB
}

func (repo *PersonRepo) Get(id string) (Person, error) {
	var m Person

	dbResult := repo.DB.First(&m, id)

	if dbResult.RecordNotFound() {
		return m, dbResult.Error
	} else {
		log.Print(dbResult)

		return m, nil
	}
}

//func (repo *PersonRepo) List(c *gin.Context) {
//var m Person

//dbResult := repo.DB.First(&m, c.Param("id"))

//if dbResult.RecordNotFound() {
//c.AbortWithStatus(404)
//} else {
//c.JSON(200, m)
//}
//}

//func (repo *PersonRepo) Create(c *gin.Context) {
//var m Person

//c.BindJSON(&m)

//dbResult := repo.DB.Create(&m)

//if dbResult.Error != nil {
//log.Print(dbResult.Error)
//c.JSON(422, dbResult.Error)
//} else {
//c.JSON(201, m)
//}
//}

//func (repo *PersonRepo) Delete(c *gin.Context) {
//}

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
