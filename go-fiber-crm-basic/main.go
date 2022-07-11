package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(getlead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
}
