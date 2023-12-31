package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/rest-go-fiber-docker/database/db"
	"github.com/joewilson27/rest-go-fiber-docker/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	db.DB.Find(&facts)

	//return c.Status(http.StatusOK).JSON(facts)

	// render template to response
	return c.Render("index", fiber.Map{
		"Title":    "Fiber Docker Trivia",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts":    facts, // bounding data facts from database
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

	result := db.DB.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}
	//return c.Status(http.StatusOK).JSON(fact)

	// return to confirmation page view
	return ListFacts(c)
}

// New Confirmation view
func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := db.DB.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Status(fiber.StatusOK).Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := db.DB.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"Title":    "Edit Fact",
		"Subtitle": "Edit your interesting fact",
		"Fact":     fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	// // Write updated values to the database
	// result := database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact)
	// if result.Error != nil {
	// 	return EditFact(c)
	// }
	jsonString, _ := json.Marshal(fact)
	var factRequest models.Fact
	json.Unmarshal(jsonString, &factRequest)
	updId, _ := strconv.Atoi(id)
	factRequest.ID = updId
	if err := factRequest.Update(); err != nil {
		return EditFact(c)
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := db.DB.Where("id = ?", id).Delete(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListFacts(c)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}
