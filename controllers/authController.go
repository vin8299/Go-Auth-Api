package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/authenticate"
)
var data map[string]interface{}

var err error

func Register(c *fiber.Ctx)(error) {
	var output=make(map[string]interface{})

	if err:=c.BodyParser(&data);err!=nil{
		return err
	}

	output = models.RegisterUser(data)

	return c.JSON(output)
}

func Login(c *fiber.Ctx) (error){
	
	var output=make(map[string]interface{})

	if err:=c.BodyParser(&data);err!=nil{
		return err
	}
		output = models.LoginUser(data,c)
	


	return c.JSON(output)
}

func UserDetail(c *fiber.Ctx) (error){
	var output=make(map[string]interface{})

	if err:=c.BodyParser(&data);err!=nil{
		return err
	}

	authenticated,_:=auth.AuthUser(data,c)
	if !authenticated{
		output["Message"]="Unauthenticated user"
        output["Status"]="failure"
		output["Code"]="401"
		output["Data"]=""

	}else{
		output = models.UserDetails(data)
	}

	return c.JSON(output)
}

func Logout(c *fiber.Ctx) (error){
	var output=make(map[string]interface{})
	//Authenticate if the same user is trying to logout
	if err:=c.BodyParser(&data);err!=nil{
		return err
	}

	authenticated,_:=auth.AuthUser(data,c)
	if !authenticated{
		output["Message"]="Not authorised to logout"
        output["Status"]="failure"
		output["Code"]="401"

	}else{
		output = models.Logout(c)
	}

	return c.JSON(output)

}
