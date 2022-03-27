package main

import (
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

    log.Fatal(app.Listen(":8000"))
}