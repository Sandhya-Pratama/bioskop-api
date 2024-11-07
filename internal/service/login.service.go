package service

import (
	"context"
	"errors"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
	"golang.org/x/crypto/bcrypt"
)

// login
type LoginUseCase interface {
	Login(ctx context.Context, username string, password string) (*entity.User, error)
}

type LoginRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type loginService struct {
	repository LoginRepository
}

func NewLoginService(repository LoginRepository) *loginService {
	return &loginService{
		repository: repository,
	}
}

// func untuk melakikan pengecekan apakah semua data nya sama mulai dari username, password
func (s *loginService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := s.repository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	//untuk pengecakan apakah username  ada di database?
	if user == nil {
		return nil, errors.New("user with that username not found")
	}

	//untuk pengecekan apakah password nya ada atau gaa di databse?
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}
	//ketika username dan passwerd sama maka akan mengembalikan nil
	return user, nil

}
