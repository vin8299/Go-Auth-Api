package routes
import(
	"github.com/gofiber/fiber/v2"
	"main/controllers"
)

func Setup(app *fiber.App){
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login",controllers.Login)
}