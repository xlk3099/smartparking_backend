package model

// Park
type Parking struct {
	ID          int    `json:"id" binding:"required"`
	Available   bool   `json:"available" binding:"required"`
	PlateNumber string `json:"plateNo" binding:"required"`
}
