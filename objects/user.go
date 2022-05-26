package objects

import (
	"time"
)

type User struct {
	ID          string    `gorm:"type:varchar(50);primaryKey;not null" json:"id" example:"服务端会自动生成"`
	Name        string    `gorm:"type:varchar(100);unique;uniqueIndex;not null" json:"name"`
	DateOfBirth time.Time `gorm:"column:dob;not null" json:"dob"`
	Address     string    `gorm:"type:varchar(256);not null" json:"address"`
	Description string    `gorm:"type:varchar(512)" json:"description"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null" json:"createdAt" example:"服务端会自动生成"`
}
