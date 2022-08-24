package repositories

import (
	"database/sql"
	"go-sqlx/models"
	"log"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return ItemRepository{db}
}

func (itemRepo ItemRepository) CountAll() int {
	log.Println("=== FETCHING ALL ITEMS ===")
	return 0
}

func (itemRepo ItemRepository) FetchAll() []models.Item {
	log.Println("=== FETCHING ALL ITEMS ===")

	return []models.Item{}
}

func (itemRepo ItemRepository) FetchById(itemId int) models.Item {
	log.Println("=== FETCH ITEM BY ID ===")

	return models.Item{}
}

func (itemRepo ItemRepository) Create(item models.Item) {
	log.Println("=== CREATE NEW ITEM ===")

}

func (itemRepo ItemRepository) UpdateById(id int, item models.Item) {
	log.Println("=== UPDATE ITEM BY ID ===")

}

func (itemRepo ItemRepository) DeleteById(id int) {
	log.Println("=== DELETE ITEM BY ID ===")
}
