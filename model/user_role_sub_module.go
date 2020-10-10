package model

type UserRoleSubModule struct {
	Model
	UserRoleId  int `json:"user_role_id"`
	SubModuleId int `json:"sub_module_id"`
}
