package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement" pg:"id_question"`
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}
