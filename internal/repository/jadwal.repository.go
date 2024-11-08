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

func (r *JadwalTayangRepository) CreateJadwal(ctx context.Context, jadwal *entity.Jadwal_Tayang) error {
	err := r.db.WithContext(ctx).Create(&jadwal).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *JadwalTayangRepository) GetJadwalByID(ctx context.Context, id int64) (*entity.Jadwal_Tayang, error) {
	jadwal := new(entity.Jadwal_Tayang)
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&jadwal).Error
	if err != nil {
		return nil, err
	}
	return jadwal, nil
}

func (r *JadwalTayangRepository) UpdateJadwal(ctx context.Context, jadwal *entity.Jadwal_Tayang) error {
	if err := r.db.WithContext(ctx).
		Model(&entity.Jadwal_Tayang{}).
		Where("jadwal_id = ?", jadwal.JadwalID).
		Updates(jadwal).Error; err != nil {
		return err
	}
	return nil
}

// DeleteJadwal menghapus jadwal tayang berdasarkan ID
func (r *JadwalTayangRepository) DeleteJadwal(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Jadwal_Tayang{}, id).Error; err != nil {
		return err
	}
	return nil
}
