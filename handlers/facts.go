package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/rest-go-fiber-docker/database"
	"github.com/joewilson27/rest-go-fiber-docker/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(http.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	var fact models.Fact
	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(http.StatusOK).JSON(fact)
}