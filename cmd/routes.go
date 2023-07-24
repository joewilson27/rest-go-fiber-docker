package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/rest-go-fiber-docker/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)

	app.Post("/fact", handlers.CreateFact)

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)

	// Display `Edit` form
	app.Get("/fact/:id/edit", handlers.EditFact)
	// Update fact
	app.Patch("/fact/:id", handlers.UpdateFact)

	// Delete fact
	app.Delete("/fact/:id", handlers.DeleteFact)
}
