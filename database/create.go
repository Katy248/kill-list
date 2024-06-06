package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() {
	db := CreateDB()

	err := db.AutoMigrate(&User{}, &ListItem{}, &ListItemDataType{}, &ListItemData{})
	if err != nil {
		log.Fatalln(err)
	}
}
func CreateDB() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	timezone := "Moscow/Europe"

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=" + timezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db

}
