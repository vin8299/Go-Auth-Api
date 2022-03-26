package main

import (
    "log"
    "main/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    routes.Setup(app)

    log.Fatal(app.Listen(":8000"))
}