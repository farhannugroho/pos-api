package model

type Item struct {
	Model
	Sku                 string  `json:"sku"`
	Barcode             string  `json:"barcode"`
	Name                string  `json:"name,omitempty"`
	Description         string  `json:"description"`
	ImageUrl            string  `json:"image_url"`
	Price               float64 `json:"price"`
	ItemGroupId         int     `json:"item_group_id"`
	ItemCategoryId      int     `json:"item_category_id"`
	OutletId            int     `json:"outlet_id"`
	UnitOfMeasurementId int     `json:"unit_of_measurement_id"`
	IsActive            bool    `json:"is_active"`
}
