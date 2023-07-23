package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/rest-go-fiber-docker/database"
	"github.com/joewilson27/rest-go-fiber-docker/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	//return c.Status(http.StatusOK).JSON(facts)

	// render template to response
	return c.Render("index", fiber.Map{
		"Title": "Fiber Docker Trivia",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts": facts, // bounding data facts from database
	})
}

// Create new Fact View handler
/*
	template: new.html
*/
func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	var fact models.Fact
	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	//return c.Status(http.StatusOK).JSON(fact)

	// return to confirmation page view
	return ConfirmationView(c)
}

// New Confirmation view
func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}