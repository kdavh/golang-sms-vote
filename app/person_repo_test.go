package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kdavh/gin-sms-vote/db"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	// we register an sql driver named "txdb"
	//txdb.Register("txdb", "postgres", "root@/txdb_test")

	m.Run()
}

func initTestDB() *gorm.DB {
	d := db.InitEnv("test")
	d.AutoMigrate(&Person{})

	return d
}

func TestPersonRepo(t *testing.T) {
	// initialize DB as transaction for easy rollback
	d := initTestDB().Begin()
	defer d.Close()

	firstName := "asdf"
	p := Person{FirstName: firstName, LastName: "ddd", Email: "sdjdkd"}

	d.Create(&p)

	repo := PersonRepo{DB: d}

	person, _ := repo.Get(strconv.Itoa(int(p.ID)))

	assert.Equal(t, person.FirstName, firstName)
}
