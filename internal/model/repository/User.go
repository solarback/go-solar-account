package repository

type User struct {
	Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string
	Password string
	Accounts []Account `gorm:"foreignKey:UserId"`
}
