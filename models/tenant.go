package models

type Tenant struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
