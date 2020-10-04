package model

type ItemGroup struct {
	Model
	Code                string `json:"code"`
	Name                string `json:"name"`
	UnitOfMeasurementId int    `json:"unit_of_measurement_id"`
	IsActive            bool   `json:"is_active"`
}
