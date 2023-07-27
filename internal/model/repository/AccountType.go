package repository

type AccountType struct {
	Type        string `gorm:"primary_key;type:varchar(32)"`
	Title       string `gorm:"type:varchar(32)"`
	Description string `gorm:"type:varchar(255)"`
	IsClient    bool
}
