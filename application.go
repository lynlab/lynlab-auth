package main

import (
	"time"

	"github.com/lib/pq"
)

type Application struct {
	ID          int64          `gorm:"primary_key"`
	UUID        string         `gorm:"type:varchar(36);unique_index"`
	Name        string         `gorm:"type:varchar(32);unique"`
	RedirectURL string         `gorm:"text"`
	Scopes      pq.StringArray `gorm:"type:varchar(10)[]"`
	OwnerID     int64          `gorm:"index"`
	CreatedAt   time.Time      `gorm:"assocation_autocreate"`
	UpdatedAt   time.Time      `gorm:"assocation_autoupdate"`
}
