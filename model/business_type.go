package model

type BusinessType struct {
	Model
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
