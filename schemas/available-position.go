package schemas

import (
	"time"

	"gorm.io/gorm"
)

type AvailablePosition struct {
	gorm.Model
	Role     string
	Tech     string
	Level    string
	Company  string
	Location string
	Salary   float64
	Link     string
}

type AvailablePositionResponse struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Tech      string    `json:"tech"`
	Level     string    `json:"level"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Salary    float64   `json:"salary"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}
