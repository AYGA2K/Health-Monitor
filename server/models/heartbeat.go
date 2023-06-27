package models

import "time"

type Heartbeat struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Value     int  `json:"value"`
	UserID    uint `json:"userId"`
}
