package service

import (
	"context"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
)

type JadwalTayangUsecase interface {
	CreateJadwal(ctx context.Context, jadwal *entity.JadwalTayang) error
}

type JadwalTayangRepository interface {
	CreateJadwal(ctx context.Context, jadwal *entity.JadwalTayang) error
}

type JadwalTayangService struct {
	repository JadwalTayangRepository
}

func NewJadwalTayangService(repository JadwalTayangRepository) *JadwalTayangService {
	return &JadwalTayangService{repository}
}

func (s *JadwalTayangService) CreateJadwal(ctx context.Context, jadwal *entity.JadwalTayang) error {
	return s.repository.CreateJadwal(ctx, jadwal)
}