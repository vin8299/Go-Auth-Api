package models

import(
	"github.com/gofiber/fiber/v2"
	"time"
)

func Logout(c *fiber.Ctx)(map[string]interface{}){

	var response map[string]interface{} = map[string]interface{}{
		"Code":"",
		"Status":"",
		"Message":"",
	}

	//Create a new cookie with past time
	cookie := fiber.Cookie{
		Name:		"jwt",
		Value:		" ",
		Expires:	time.Now().Add(-time.Hour),
		HTTPOnly:	true,//Front-end will not be able to access it,it will be only stored and will be sent 
	}
	c.Cookie(&cookie)

	response["Code"]="200"
	response["Message"]="Logged out successfully"
	response["Status"]="success"

	return response
}