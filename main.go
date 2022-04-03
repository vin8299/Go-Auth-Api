package main

import (
    "os"
    "log"
    "main/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()
 
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
    }))
    routes.Setup(app)
    port := os.Getenv("PORT")
    if port == "" {
          port = "8080"
  }
    log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}