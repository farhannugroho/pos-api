package model

type Module struct {
	Model
	Name      string      `json:"name"`
	ModuleId  int         `gorm:"-" json:"module_id"`
	SubModule []SubModule `gorm:"foreignKey:ModuleId" json:"sub_module"`
}
