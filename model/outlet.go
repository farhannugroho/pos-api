package model

type Outlet struct {
	Model
	Code       string   `json:"code"`
	Name       string   `json:"name"`
	ImageUrl   string   `json:"image_url"`
	Address    string   `json:"address"`
	Phone      string   `json:"phone"`
	CityId     int      `json:"city_id"`
	City       City     `gorm:"foreignKey:CityId" json:"city"`
	LocationId int      `json:"location_id"`
	Location   Location `gorm:"foreignKey:LocationId" json:"location"`
	CompanyId  int      `json:"company_id"`
	IsActive   bool     `json:"is_active"`
}
