package main

import (
	"os"
	"personal-financial-tracker/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var HOST string
var MYSQL_DATABASE_URL string
var MYSQL_DSN string

func LoadEnv() {
	// err := godotenv.Load(path + ".env")
	// if err != nil {
	// 	logger.Log.Error(err)
	// }

	// var SECRET string
	HOST = os.Getenv("HOST")
	MYSQL_DATABASE_URL = os.Getenv("MYSQL_DATABASE_URL")
	MYSQL_DSN = os.Getenv("MYSQL_DSN")
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 500 * 1024 * 1024, // 500 MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	Initialize(app)

	err := app.Listen(HOST)
	if err != nil {
		panic("Failed to listen to host: " + err.Error())
	}
}

func Initialize(app *fiber.App) {
	LoadEnv()

	var gdb *gorm.DB

	gdb = ConnectDB()
	// MigrateDB(gdb)

	router.SetupRoutes(app, gdb)
}

func ConnectDB() *gorm.DB {

	gdb, err := gorm.Open(mysql.Open(MYSQL_DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	return gdb
}

// func MigrateDB(gdb *gorm.DB) {
// 	err := gdb.AutoMigrate(
// 		model.Transaction{},
// 	)
// 	if err != nil {
// 		panic("Failed to migrate database" + err.Error())
// 	}
// }
