package model

type Item struct {
	Model
	Name                string        `json:"name,omitempty"`
	Description         string        `json:"description"`
	ImageUrl            string        `json:"image_url"`
	ItemGroupId         int           `json:"item_group_id"`
	ItemCategoryId      int           `json:"item_category_id"`
	OutletId            int           `json:"outlet_id"`
	UnitOfMeasurementId int           `json:"unit_of_measurement_id"`
	ItemId              int           ` gorm:"-" json:"item_id"`
	ItemVariant         []ItemVariant `gorm:"foreignKey:ItemId" json:"item_variant"`
	IsActive            bool          `json:"is_active"`
}
