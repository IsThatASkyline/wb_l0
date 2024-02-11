package postgres

import (
	"github.com/IsThatASkyline/wb_l0/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func GetItems(db *sqlx.DB, track_number string) ([]models.Item, error) {
	items := make([]models.Item, 0)
	query := `SELECT * FROM items WHERE "track_number" = $1`
	err := db.Select(&items, query, track_number)

	return items, err
}

func SetItems(items []models.Item, db *sqlx.DB) {
	query := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	for _, item := range items {
		_, err := db.Exec(query, item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale,
			item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status)
		if err != nil {
			log.Print("error SetPayment postgres")
		}
	}
}
