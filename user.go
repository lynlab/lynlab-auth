package main

import (
	"time"

	"github.com/lib/pq"
	"github.com/thanhpk/randstr"
)

type UserIdentity struct {
	ID        int64  `gorm:"primary_key"`
	UUID      string `gorm:"varchar(36);unique_index"`
	Username  string `gorm:"type:varchar(40);unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (id *UserIdentity) NewToken(scopes []string) *UserToken {
	token := UserToken{
		AccessToken:          randstr.String(80),
		RefreshToken:         randstr.String(80),
		Scopes:               scopes,
		IdentityID:           id.ID,
		AccessTokenExpireAt:  time.Now().AddDate(0, 0, 7),
		RefreshTokenExpireAt: time.Now().AddDate(0, 0, 30),
	}
	db.Save(&token)
	return &token
}

type UserAccount struct {
	ID               int64  `gorm:"primary_key"`
	IdentityID       int64  `gorm:"index"`
	Provider         string `gorm:"type:varchar(20);index:idx_provider_provider_identity"`
	ProviderIdentity string `gorm:"type:text;index:idx_provider_provider_identity"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (a *UserAccount) GetIdentity() *UserIdentity {
	var i UserIdentity
	db.Where(&UserIdentity{ID: a.IdentityID}).First(&i)
	return &i
}

type UserToken struct {
	ID                   int64          `gorm:"primary_key"`
	AccessToken          string         `gorm:"type:varchar(80);unique_index"`
	RefreshToken         string         `gorm:"type:varchar(80);unique_index"`
	Scopes               pq.StringArray `gorm:"type:varchar(10)[]"`
	IdentityID           int64          `gorm:"index"`
	AccessTokenExpireAt  time.Time
	RefreshTokenExpireAt time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (t *UserToken) Expired() bool {
	return t.AccessTokenExpireAt.Before(time.Now())
}

func (t *UserToken) GetIdentity() *UserIdentity {
	var i UserIdentity
	db.Where(&UserIdentity{ID: t.IdentityID}).First(&i)
	return &i
}

type UserAllowedApplication struct {
	ID            int64 `gorm:"primary_key"`
	ApplicationID int64 `gorm:"index"`
	IdentityID    int64 `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
