// model/user.go
package model

import "gorm.io/gorm"

// User
type User struct {
    gorm.Model
    ID       string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid();not null"`
    Name     string `json:"name"`
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `json:"password"`
    Phone    string `gorm:"unique; not null" json:"phone"`
}

// LoginInput
type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
