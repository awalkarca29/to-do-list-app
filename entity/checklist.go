package entity

import "time"

type Checklist struct {
	ID        int
	Title     string
	UserID    int
	Users     User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
