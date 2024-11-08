package repository

import (
	"context"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
	"gorm.io/gorm"
)

type JadwalTayangRepository struct {
	db *gorm.DB
}

func NewJadwalTayangRepository(db *gorm.DB) *JadwalTayangRepository {
	return &JadwalTayangRepository{db: db}
}

func (r *JadwalTayangRepository) CreateJadwal(ctx context.Context, jadwal *entity.JadwalTayang) error {
	err := r.db.WithContext(ctx).Create(&jadwal).Error
	if err != nil {
		return err
	}
	return nil
}
