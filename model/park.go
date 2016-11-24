package model

// Park
type Parking struct {
	Id     int    `json:"id"`
	Number string `json:"number"`
	Empty  bool   `json:"empty"`
}
