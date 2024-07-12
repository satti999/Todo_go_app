package model

type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	Content   *string   `json:"content"`
	Completed *bool     `json:"completed"`
	Createdby uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:Createdby;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubTodos  []SubTodo `gorm:"foreignKey:TodoID"`
}
type SubTodo struct {
	ID        uint    `gorm:"primaryKey"`
	TodoID    uint    `gorm:"not null"`
	Content   *string `json:"content"`
	Completed *bool   `json:"completed"`
}

// type Todo struct {
// 	gorm.Model
// 	ID        uint    `gorm:"autoIncrement" `
// 	Content   *string `json:"content"`
// 	Completed *bool   `json:"completed"`
// 	Createdby uint

// 	User     User `gorm:"foreignKey:Createdby references:UserID"`
// 	TodoID   uint
// 	SubTodos []SubTodo `gorm:"foreignKey:TodoID"`
// }

// type SubTodo struct {
// 	gorm.Model
// 	TodoID    uint    `gorm:"autoIncrement"`
// 	Content   *string `json:"content"`
// 	Completed *bool   `json:"completed"`
// 	Createdby uint

// 	User User `gorm:"foreignKey:Createdby references:UserID"`
// }
