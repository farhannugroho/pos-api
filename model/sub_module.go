package model

type SubModule struct {
	Model
	Name     string `json:"name"`
	ModuleId int    `json:"module_id"`
}
