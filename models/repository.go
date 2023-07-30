package models

import (
	"github.com/joewilson27/rest-go-fiber-docker/database/db"
)

func (fact Fact) Update() error {
	result := db.DB.Model(&fact).Updates(fact)
	return result.Error
}
