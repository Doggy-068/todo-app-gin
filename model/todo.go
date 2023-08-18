package model

type Todo struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"not null"`
	IsComplete bool   `gorm:"not null"`
}
