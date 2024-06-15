package service

import (
	"gpt/pkg/repository"
)

type Authorization interface {
}


type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
