package cache

import (
	"errors"
	"github.com/IsThatASkyline/wb_l0/models"
	"sync"
)

type Cache struct {
	sync.RWMutex
	orders map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{
		orders: make(map[string]models.Order),
	}
}

func (c *Cache) Set(uid string, order models.Order) {
	c.Lock()
	defer c.Unlock()
	c.orders[uid] = order
}

func (c *Cache) GetOrderByUid(uid string) (models.Order, error) {
	c.RLock()
	defer c.RUnlock()
	order, ok := c.orders[uid]
	if !ok {
		return order, errors.New("error cache for GetOrderByUid")
	}
	return order, nil
}
