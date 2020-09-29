package model

type Outlet struct {
	Model
	Code       string `json:"code"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	CityId     int    `json:"city_id"`
	LocationId int    `json:"location_id"`
	CompanyId  int    `json:"company_id"`
	IsActive   bool   `json:"is_active"`
}
