package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Tods     []Todo `gorm:"foreignKey:Createdby;references:ID"`
}
