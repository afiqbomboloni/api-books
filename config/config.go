package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Config struct {
	DBDriver string
	DBHost   string
	DBPass   string
	DBName   string
	DBUser   string
	DBPort   int
}

func LoadConfig() (Config) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file failed")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	return Config{
		DBDriver: os.Getenv("DB_DRIVER"),
		DBHost: os.Getenv("DB_HOST"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPort: dbPort,
	}
}


func ConnectDb() *gorm.DB{

	cfg := LoadConfig()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db

}