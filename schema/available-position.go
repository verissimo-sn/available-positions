package schema

import (
	"gorm.io/gorm"
)

type AvailablePosition struct {
	gorm.Model
	Role     string
	Techs    []string
	Level    string
	Company  string
	Location string
	Salary   float64
	Link     string
}
