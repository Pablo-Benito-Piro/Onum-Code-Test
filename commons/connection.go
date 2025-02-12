package commons

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"onumTest/models"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/onumtest?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitialMigrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migrating database...")
	db.AutoMigrate(&models.Bid{}, &models.Auction{})
}
