package entity

import (
	"time"
)

type User struct {
	User_ID   int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string     `json:"name"`
	Password  string     `json:"-"`
	Roles     string     `json:"roles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Admin New User
func NewUser(username, password, roles string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Roles:     roles,
		CreatedAt: time.Now(),
	}
}

// Admin Update User
func UpdateUser(user_id int64, username, roles, password string) *User {
	return &User{
		User_ID:   user_id,
		Username:  username,
		Password:  password,
		Roles:     roles,
		UpdatedAt: time.Now(),
	}
}

// Public Register
func Register(username, password, roles string) *User {
	return &User{
		Username: username,
		Password: password,
		Roles:    roles,
	}
}

// Public Login
func Login(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

func (User) TableName() string {
	return "User" // Pastikan sesuai dengan nama tabel di PostgreSQL
}
