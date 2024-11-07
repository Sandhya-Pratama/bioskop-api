package repository

import (
	"context"
	"errors"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
	"gorm.io/gorm"
)

const (
	Userkey = "users:all"
)

type UserRepository struct {
	db *gorm.DB
}

// membuat constructor untuk dependency
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db}
}

// menampilkan get all user
// menggunakan []*entity.User = karena akan membutuhkan data yg banyak dengan array slice of user.
func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	// Inisialisasi slice untuk menyimpan data pengguna
	users := make([]*entity.User, 0)

	// Melakukan query ke database untuk mendapatkan semua pengguna
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// membuat create user
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	//menggunakan db untuk melakukan query ke database
	err := r.db.WithContext(ctx).Create(&user).Error // pada line ini akan melakukan query "INSERT INTO users"
	if err != nil {
		return err
	}
	return nil
}

// update data user byID
func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.User_ID).
		Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// get user by id
func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// detele user by id
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// get user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("username not found")
	}
	return user, nil
}
