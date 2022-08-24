package repositories

import (
	"database/sql"
	"go-sqlx/models"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{db}
}

func (orderRepo OrderRepository) CreateOrder(order models.Order, item models.Item, orderQty int) {
	log.Println("=== CREATE ORDER ===")
}
