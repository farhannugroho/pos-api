package model

type UserRole struct {
	Model
	Code      string      `json:"code"`
	Name      string      `json:"name"`
	CompanyId int         `json:"company_id"`
	Scopes    []SubModule `gorm:"many2many:userRole_subModules;"`
	IsActive  bool        `json:"is_active"`
}
