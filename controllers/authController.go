package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
)
var data map[string]interface{}
var err error

func Register(c *fiber.Ctx)(error) {

	if err:=c.BodyParser(&data);err!=nil{
		return err
	}

	data = models.RegisterUser(data)

	return c.JSON(data)
}

func Login(c *fiber.Ctx) (error){
	
	if err:=c.BodyParser(&data);err!=nil{
		return err
	}

	data = models.LoginUser(data)

	return c.JSON(data)
}