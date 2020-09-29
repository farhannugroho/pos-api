package model

type Company struct {
	Model
	Code           string `json:"code"`
	Name           string `json:"name"`
	ImageUrl       string `json:"image_url"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Fax            int    `json:"fax"`
	Email          string `json:"email"`
	CityId         int    `json:"city_id"`
	LocationId     int    `json:"location_id"`
	BusinessTypeId int    `json:"business_type_id"`
}
