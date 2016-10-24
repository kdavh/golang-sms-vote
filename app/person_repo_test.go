package app

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	if matched, _ := regexp.MatchString("_test$", os.Getenv("DB_NAME")); !matched {
		log.Fatal("must be using test database, check config")
	}

	os.Exit(m.Run())
}

func TestPersonRepoCreate(t *testing.T) {
	// initialize DB as transaction for easy rollback
	d := InitDB().Begin()
	defer d.Close()

	firstName := "asdf"
	p := Person{FirstName: firstName, LastName: "ddd", Email: "sdjdkd"}

	repo := PersonRepo{DB: d}

	err := repo.Create(&p)

	assert.Equal(t, d.NewRecord(p), false)
	assert.Nil(t, err)
}

func TestPersonRepoGet(t *testing.T) {
	// initialize DB as transaction for easy rollback
	d := InitDB().Begin()
	defer d.Close()

	firstName := "asdf"
	p := Person{FirstName: firstName, LastName: "ddd", Email: "sdjdkd"}

	d.Create(&p)

	repo := PersonRepo{DB: d}

	person, _ := repo.Get(strconv.Itoa(int(p.ID)))

	assert.Equal(t, person.FirstName, firstName)
}
