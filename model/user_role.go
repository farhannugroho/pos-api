package model

type UserRole struct {
	Model
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	CompanyId int         `json:"company_id"`
	Scopes    []SubModule `gorm:"many2many:UserRole_SubModules;"`
	IsActive  bool        `json:"is_active"`
}
