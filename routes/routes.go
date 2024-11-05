package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/briannkhata/eschool/controllers"
)

func Setup(app *fiber.App) {
	// Define cashier routes
	app.Post("/cashiers/", controllers.CreateCashier)
	app.Get("/cashiers/", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Post("/cashiers/:cashierId", controllers.UpdateCashier)

	// Uncomment these lines if you implement these controllers later
	// app.Post("/cashiers/:cashierId/login", controllers.Login)
	// app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	// app.Post("/cashiers/:cashierId/passcode", controllers.Passcode)
}
