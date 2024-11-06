package models

import "time"

type Address struct {
	ID         string `gorm:"size:36;not null;uniqeIndex;primary_key"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Name       string `gorm:"size:100"`
	IsPrimary  bool
	CityId     string `gorm:"size;100"`
	ProvinceId string `gorm:"size;100"`
	Address1   string `gorm:"size;255"`
	Address2   string `gorm:"size;255"`
	Phone      string `gorm:"size;100"`
	Email      string `gorm:"size;100"`
	PostCode   string `gorm:"size;100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
