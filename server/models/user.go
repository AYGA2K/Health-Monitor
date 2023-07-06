package models

import "time"

type User struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string      `json:"name"`
	Email        string      `json:"email" gorm:"unique"`
	Mac          string      `json:"mac"`
	Address			 string			 `json:"address"`
	Password     string      `json:"password"`
	Tokens       []Token     `json:"tokens"`
	Heartbeats   []Heartbeat `json:"hearbeats"`
	Steps        []Step      `json:"steps"`
	SleepingTime []Sleep     `json:"sleeping_time"`
}
