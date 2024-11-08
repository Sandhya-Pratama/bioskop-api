package service

import (
	"context"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
)

type JadwalTayangUsecase interface {
	CreateJadwal(ctx context.Context, jadwal *entity.Jadwal_Tayang) error
	GetJadwalByID(ctx context.Context, id int64) (*entity.Jadwal_Tayang, error)
	UpdateJadwal(ctx context.Context, user *entity.Jadwal_Tayang) error
	DeleteJadwal(ctx context.Context, id int64) error
}

type JadwalTayangRepository interface {
	CreateJadwal(ctx context.Context, jadwal *entity.Jadwal_Tayang) error
	GetJadwalByID(ctx context.Context, id int64) (*entity.Jadwal_Tayang, error)
	UpdateJadwal(ctx context.Context, user *entity.Jadwal_Tayang) error
	DeleteJadwal(ctx context.Context, id int64) error
}

type JadwalTayangService struct {
	repository JadwalTayangRepository
}

func NewJadwalTayangService(repository JadwalTayangRepository) *JadwalTayangService {
	return &JadwalTayangService{repository}
}

func (s *JadwalTayangService) CreateJadwal(ctx context.Context, jadwal *entity.Jadwal_Tayang) error {
	return s.repository.CreateJadwal(ctx, jadwal)
}

func (s *JadwalTayangService) GetJadwalByID(ctx context.Context, id int64) (*entity.Jadwal_Tayang, error) {
	return s.repository.GetJadwalByID(ctx, id)
}

func (s *JadwalTayangService) UpdateJadwal(ctx context.Context, user *entity.Jadwal_Tayang) error {
	return s.repository.UpdateJadwal(ctx, user)
}

func (s *JadwalTayangService) DeleteJadwal(ctx context.Context, id int64) error {
	return s.repository.DeleteJadwal(ctx, id)
}
