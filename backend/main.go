package main

import(
    "github.com/gofiber/fiber/v2"
	"github.com/Joe-BN/todoapp/database"
	"github.com/Joe-BN/todoapp/router"  
)

func main() {
    app := fiber.New()

    // Connect to the Database
    database.ConnectDB()

    // Setup the router
    router.SetupRoutes(app)

    // Listen on PORT 3000
    app.Listen(":3000")
}