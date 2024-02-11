package service

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/IsThatASkyline/wb_l0/pkg/repository"
)

type Order interface {
	GetOrderByUid(uid string) (models.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
