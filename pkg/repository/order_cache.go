package repository

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/IsThatASkyline/wb_l0/pkg/repository/cache"
)

type OrderCache struct {
	cache *cache.Cache
}

func NewOrderCache(cache *cache.Cache) *OrderCache {
	return &OrderCache{cache: cache}
}

func (r *OrderCache) GetOrderByUid(uid string) (models.Order, error) {
	order, err := r.cache.GetOrderByUid(uid)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
