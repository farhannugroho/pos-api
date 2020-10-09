package model

type ItemVariant struct {
	Model
	Sku      string  `json:"sku"`
	Barcode  string  `json:"barcode"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ItemId   int     `json:"item_id"`
	IsActive bool    `json:"is_active"`
}
