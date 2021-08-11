package Config

import (
	"fmt"
	"os"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

type DBConfig struct{
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig{

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("env file didn't load properly")
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}