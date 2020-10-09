package model

type Inventory struct {
	Model
	OpeningBalance int `json:"opening_balance"`
	StockIn        int `json:"stock_in"`
	StockOut       int `json:"stock_out"`
	SalesOut       int `json:"sales_out"`
	Adjustment     int `json:"adjustment"`
	EndingBalance  int `json:"ending_balance"`
	ItemVariantId  int `json:"item_variant_id"`
	OutletId       int `json:"outlet_id"`
}
