package analysis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Review struct {
	gorm.Model
	Review    string
	WordCount int
	Sentiment string
	Score     int
}
type Image struct {
	gorm.Model
	ImageURL       string
	ProminentColor string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
func ConnectToDB() {
	ipAddress := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		panic("failed to parse port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", ipAddress, username, password, dbName, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db = database
	if err = db.Exec("DELETE FROM reviews").Error; err != nil {
		panic(err)
	}
	if err = db.Exec("DELETE FROM images").Error; err != nil {
		panic(err)
	}
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Image{})
}
