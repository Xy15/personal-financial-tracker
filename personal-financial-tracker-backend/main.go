package main

import (
	"os"
	"personal-financial-tracker/router"

	_ "personal-financial-tracker/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var HOST string

// var MYSQL_DATABASE_URL string
var MYSQL_DSN string

func LoadEnv() {
	HOST = os.Getenv("HOST")
	// MYSQL_DATABASE_URL = os.Getenv("MYSQL_DATABASE_URL")
	MYSQL_DSN = os.Getenv("MYSQL_DSN")
}

// @title Personal Financial Tracker API
// @version 2.0
// @description This is a swagger for Personal Financial Tracker API
// @termsOfService http://swagger.io/terms/
// @contact.name xy
// @contact.email chongxy1115@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8081
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @BasePath /
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

	router.SetupRoutes(app, gdb)
}

func ConnectDB() *gorm.DB {

	gdb, err := gorm.Open(mysql.Open(MYSQL_DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	return gdb
}
