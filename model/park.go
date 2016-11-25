package model

// Park
type Parking struct {
	ID          int    `json:"id"`
	Available   bool   `json:"available"`
	PlateNumber string `json:"plateNo"`
}
