package service

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/IsThatASkyline/wb_l0/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrderByUid(uid string) (models.Order, error) {
	return s.repo.GetOrderByUid(uid)
}
