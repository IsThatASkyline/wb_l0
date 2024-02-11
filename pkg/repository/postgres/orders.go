package postgres

import (
	"errors"
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func SetOrder(order models.Order, db *sqlx.DB) {
	query := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, 
		delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(query, order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature,
		order.Customer_id, order.Delivery_service, order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard)
	if err != nil {
		log.Print("error SetOrder postgres")
	}
	SetDelivery(order.Delivery, db, order.Order_uid)
	SetPayment(order.Payment, db)
	SetItems(order.Items, db)
}

func GetOrder(db *sqlx.DB, uid string) (models.Order, error) {
	var order models.Order
	query := `SELECT * FROM orders WHERE "order_uid" = $1`
	err := db.Get(&order, query, uid)
	if err != nil {
		return order, errors.New("the result set is empty, Orders")
	}

	return order, err
}

func GetOrders(db *sqlx.DB) ([]models.Order, error) {
	orders := make([]models.Order, 1)
	query := `SELECT * FROM orders`
	rows, err := db.Query(query)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale, &order.Internal_signature, &order.Customer_id,
			&order.Delivery_service, &order.Shardkey, &order.Sm_id, &order.Date_created, &order.Oof_shard)
		if err != nil {
			return orders, err
		}
		order.Delivery, err = GetDelivery(db, order.Order_uid)
		if err != nil {
			return orders, err
		}
		order.Payment, err = GetPayment(db, order.Order_uid)
		if err != nil {
			return orders, err
		}
		order.Items, err = GetItems(db, order.Track_number)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, err
}
