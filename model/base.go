package model

import (
	"time"
)

// Base is base of model
type Base struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:ctime" json:"ctime"`
	UpdatedAt time.Time `gorm:"column:mtime" json:"mtime"`
	DeleteAt  time.Time `gorm:"column:dtime" json:"dtime"`
}
