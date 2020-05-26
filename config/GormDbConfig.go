package Config

import (
	"fmt"

	"../src/api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

const (
	dbHost   = "35.187.246.42"
	username = "postgres"
	password = "LCSfms45FVmb"
	dbName   = "taperadb"
)

func init() {

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUrl)

	conn, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		fmt.Print(err)
	} else {

		fmt.Print("Database connected")
	}

	conn.Debug().AutoMigrate(models.Profile{})

	//conn.Create(&m)

	db = conn

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
