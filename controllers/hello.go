package controllers

import (
        "github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx)(error) {
	ouput := make(map[string]string)
        output["msg"] = "Hello World"
	return c.JSON(output)
}
