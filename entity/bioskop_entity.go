package entity

import (
	"time"
)

type Bioskop struct {
	ID           int64      `json:"id"`
	Nama_Bioskop string     `json:"nama_bioskop"`
	Lokasi       string     `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// Admin New User
func NewBioskop(username, password, roles string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Roles:     roles,
		CreatedAt: time.Now(),
	}
}

// Admin Update User
func UpdateBioskop(id int64, username, roles, password string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Roles:     roles,
		UpdatedAt: time.Now(),
	}
}
