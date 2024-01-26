package handler

type CretePositionDto struct {
	Role     string  `json:"role"`
	Tech     string  `json:"tech"`
	Level    string  `json:"level"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Salary   float64 `json:"salary"`
	Link     string  `json:"link"`
}

func (dto *CretePositionDto) Validate() error {
	return nil
}
