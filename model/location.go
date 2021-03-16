package model

type Location struct {
	Model
	Code     string `json:"code"`
	Name     string `json:"name"`
	CityId   int    `json:"city_id"`
	IsActive bool   `json:"is_active"`
}
