package database

import (
	"fmt"
	"log"
	"marketplace-gin/models"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func isDatabaseEmpty() bool {
	var count int64
	DB.Raw("SELECT COUNT(*) FROM users").Count(&count)
	return count == 0
}

func executeSQLFile(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Could not read SQL file: %v", err)
	}

	commands := strings.Split(string(content), ";")
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command != "" {
			if err := DB.Exec(command).Error; err != nil {
				log.Printf("Error executing command: %s\nError: %v", command, err)
			}
		}
	}
}

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbURL := os.Getenv("DB_URL")

	// Create database if using MySQL/MariaDB
	if dbConnection == "mysql" || dbConnection == "mariadb" {
		// Connect to MySQL without database selected
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
			dbUser, dbPass, dbHost, dbPort)
		tempDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Could not connect to MySQL server: %v", err)
		}
		// Create database if not exists
		tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	}

	var dialector gorm.Dialector

	switch dbConnection {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(fmt.Sprintf("%s.db", dbName))
	case "mysql", "mariadb":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)
		dialector = mysql.Open(dsn)
	default:
		log.Fatal("Unsupported database connection type")
	}

	// Set up logger configuration based on environment
	logLevel := logger.Silent
	if os.Getenv("GIN_MODE") != "release" {
		logLevel = logger.Error // Only log errors in development
	}

	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Create database if using MySQL/MariaDB
	if dbConnection == "mysql" || dbConnection == "mariadb" {
		// Create database if not exists
		tempDB, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s):%s/",
			dbUser, dbPass, dbHost, dbPort, dbURL)), &gorm.Config{})
		if err == nil {
			tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
		}
	}
}

func InitializeDB() {
	MigrateDB()

	// Check if database is empty before initialization
	if isDatabaseEmpty() {
		log.Println("Database is empty. Initializing with init.sql...")
		executeSQLFile("init.sql")
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{}, &models.Cart{}, &models.Category{}, &models.Product{}, &models.CartItem{}, &models.Discussion{}, &models.Order{}, &models.OrderItem{}, &models.Reply{}, &models.Review{}, &models.Transaction{}, &models.Wishlist{})
}
