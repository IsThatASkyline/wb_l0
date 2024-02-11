package controllers

import (
	"encoding/json"
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/IsThatASkyline/wb_l0/pkg/repository/cache"
	"github.com/IsThatASkyline/wb_l0/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"log"
)

func MsgHandler(cache *cache.Cache, db *sqlx.DB) func(m *stan.Msg) {
	return func(m *stan.Msg) {
		jsonString := string(m.Data)
		var order models.Order
		err := json.Unmarshal([]byte(jsonString), &order)
		if err != nil {
			log.Fatalf("error json unmarshl :%s", err)
		} else {
			ProcessCache(cache, order, db)
		}
	}
}

func ProcessCache(cache *cache.Cache, order models.Order, db *sqlx.DB) {
	if _, err := cache.GetOrderByUid(order.Order_uid); err != nil {
		cache.Set(order.Order_uid, order)
		postgres.SetOrder(order, db)
	} else {
		_, err := postgres.GetOrder(db, order.Order_uid)
		if err != nil {
			log.Fatalf("error  postgres.GetOrder in controllers CheckCache :%s", err)
		}
	}
}
