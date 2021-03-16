package model

type ItemCategory struct {
	Model
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	ItemGroupId int       `json:"item_group_id"`
	ItemGroup   ItemGroup `gorm:"foreignKey:ItemGroupId" json:"item_group"`
	IsActive    bool      `json:"is_active"`
	//Items       []Item `gorm:"foreignKey:ItemCategoryId" json:"items"`
}
