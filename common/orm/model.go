package orm

import "gorm.io/gorm"

type Model struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime      `json:"createdAt"`
	UpdatedAt LocalTime      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

type Base struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime `json:"createdAt"`
	UpdatedAt LocalTime `json:"updatedAt"`
}

type SoftDelete struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
