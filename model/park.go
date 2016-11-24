package model

// Park
type Parking struct {
	Available   bool   `json:"empty"`
	PlateNumber string `json:"number"`
}
