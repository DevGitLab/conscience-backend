package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	Level     int
}

func (User) TableName() string {
	return "user"
}
