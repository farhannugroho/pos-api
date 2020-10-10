package model

type UserRole struct {
	Model
	Code      string `json:"code"`
	Name      string `json:"name"`
	CompanyId int    `json:"company_id"`
	Scopes    []int  `gorm:"-" json:"scopes"`
	IsActive  bool   `json:"is_active"`
}
