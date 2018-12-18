package main

import "time"

type User struct {
	UUID         string `gorm:"type:varchar(40);pramiry_key"`
	Email        string `gorm:"type:varchar(255);not null;unique_index"`
	Username     string `gorm:"type:varchar(255);not null"`
	PasswordHash []byte `gorm:"not null" json:"-"`
	PasswordSalt []byte `gorm:"not null" json:"-"`
	IsActivated  bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
