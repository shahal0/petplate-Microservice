package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName   string `gorm:"not null" json:"first_name"`
    LastName    string `gorm:"not null" json:"last_name"`
    Email       string `gorm:"unique;not null" json:"email"`
    Password    string `gorm:"not null" json:"password"`
    Phone       string `gorm:"unique" json:"phone"`
    IsBlocked  bool   `gorm:"default:false" json:"is_blocked"` 
}
