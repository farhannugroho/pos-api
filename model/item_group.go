package model

type ItemGroup struct {
	Model
	Code                string            `json:"code" binding:"required"`
	Name                string            `json:"name"`
	UnitOfMeasurementId int               `json:"unit_of_measurement_id"`
	UnitOfMeasurement   UnitOfMeasurement `gorm:"foreignKey:UnitOfMeasurementId" json:"unit_of_measurement"`
	IsActive            bool              `json:"is_active"`
}
