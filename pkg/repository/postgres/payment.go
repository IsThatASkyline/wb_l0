package postgres

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func GetPayment(db *sqlx.DB, uid string) (models.Payment, error) {
	var payment models.Payment
	query := `SELECT * FROM payments WHERE "transaction" = $1`
	err := db.Get(&payment, query, uid)

	return payment, err
}

func SetPayment(payment models.Payment, db *sqlx.DB) {
	query := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Exec(query, payment.Transaction, payment.Request_id, payment.Currency, payment.Provider, payment.Amount, payment.Payment_dt,
		payment.Bank, payment.Delivery_cost, payment.Goods_total, payment.Custom_fee)
	if err != nil {
		log.Print("error SetPayment postgres")
	}
}
