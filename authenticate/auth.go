package auth

import(
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
)
var err error

var secretKey="qwert125"
func AuthUser(data map[string]interface{}, c *fiber.Ctx) (bool,error){
	cookie := c.Cookies("jwt")
	token,err:=jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func (token *jwt.Token) (interface{},error){
		return []byte(secretKey),nil
	})

	if err!=nil{
		return false,err
	}else{
		claims := token.Claims.(*jwt.StandardClaims)
		if data["id"]==claims.Issuer{
			return true,nil
		}else{
			return false,nil
		}
	}
}