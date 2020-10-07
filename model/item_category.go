package model

type ItemCategory struct {
	Model
	Code        string `json:"code"`
	Name        string `json:"name"`
	ItemGroupId int    `json:"item_group_id"`
	IsActive    bool   `json:"is_active"`
}
