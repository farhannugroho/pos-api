package model

type UnitOfMeasurement struct {
	Model
	Code      string `json:"code"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CompanyId int    `json:"company_id"`
}
