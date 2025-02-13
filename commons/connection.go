package commons

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"onumTest/models"
	"os"
)

func GetConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading env file")
	}
	userDB := os.Getenv("DATABASE_USER")
	passwordDB := os.Getenv("DATABASE_PASSWORD")
	nameDB := os.Getenv("DATABASE_NAME")
	hostDB := os.Getenv("DATABASE_HOST")
	portDB := os.Getenv("DATABASE_PORT")
	typeDB := os.Getenv("DATABASE_TYPE")
	db, err := gorm.Open(typeDB, userDB+":"+passwordDB+"@tcp("+hostDB+":"+portDB+")/"+nameDB+"?charset=utf8&parseTime=True")

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
