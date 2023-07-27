package repository

type Plan struct {
	Id          string  `gorm:"primary_key;type:varchar(32)"`
	Name        string  `gorm:"index:idx_plan_name,unique;type:varchar(255);not null"`
	Description string  `gorm:"type:varchar(255)"`
	Price       float32 `gorm:"type:float;not null"`
	Shared      bool
	Active      bool
}
