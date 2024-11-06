package routes

import (
	"github.com/briannkhata/eschool/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Define cashier routes
	app.Post("/cashiers/", controllers.CreateCashier)
	app.Get("/cashiers/", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)

	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	app.Post("/cashiers/:cashierId/passcode", controllers.Passcode)
}
