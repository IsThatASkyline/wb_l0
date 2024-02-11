package postgres

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func GetDelivery(db *sqlx.DB, uid string) (models.Delivery, error) {
	var delivery models.Delivery
	query := `SELECT * FROM deliveries WHERE "order_uid" = $1`
	err := db.Get(&delivery, query, uid)
	return delivery, err
}

func SetDelivery(delivery models.Delivery, db *sqlx.DB, order_uid string) {
	query := `INSERT INTO deliveries (name, phone, zip, city, address, region, email, order_uid) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(query, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, order_uid)
	if err != nil {
		log.Print("error SetDelivery postgres")
	}
}
