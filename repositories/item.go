package repositories

import (
	"go-sqlx/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) ItemRepository {
	return ItemRepository{db}
}

func (itemRepo ItemRepository) CountAll() int {
	log.Println("=== FETCHING ALL ITEMS ===")

	var itemCount int
	err := itemRepo.db.QueryRow(`SELECT count(*) FROM items`).Scan(&itemCount)
	if err != nil {
		log.Fatal(err)
	}
	return itemCount
}

func (itemRepo ItemRepository) FetchAll() []models.Item {
	log.Println("=== FETCHING ALL ITEMS ===")

	items := []models.Item{}

	err := itemRepo.db.Select(&items, `SELECT * FROM items`)
	if err != nil {
		log.Fatal(err)
	}

	return items
}

func (itemRepo ItemRepository) FetchById(itemId int) models.Item {
	log.Println("=== FETCH ITEM BY ID ===")

	item := models.Item{}

	err := itemRepo.db.Get(&item, `SELECT * FROM items WHERE id = $1`, itemId)
	if err != nil {
		log.Fatal(err)
	}

	return item
}

func (itemRepo ItemRepository) Create(item models.Item) {
	log.Println("=== CREATE NEW ITEM ===")

	_, err := itemRepo.db.Exec(`
		INSERT INTO items(name, qty, weight) 
		VALUES
			($1, $2, $3) 
		RETURNING id`, item.Name, item.Qty, item.Weight)

	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepo ItemRepository) CreateBatch(items []models.Item) {
	log.Println("=== CREATE NEW ITEM ===")

	_, err := itemRepo.db.NamedExec(`
		INSERT INTO items (name, qty, weight)
		VALUES (:name, :qty, :weight)`, items)

	if err != nil {
		log.Fatal(err)
	}
}

func (itemRepo ItemRepository) UpdateById(id int, item models.Item) {
	log.Println("=== UPDATE ITEM BY ID ===")

	_, err := itemRepo.db.Exec(`
		UPDATE items
		SET
			name = $1,
			qty = $2,
			weight = $3
		WHERE
			id = $4`, item.Name, item.Qty, item.Weight, item.ID)

	if err != nil {
		log.Fatal(err)
	}

}

func (itemRepo ItemRepository) DeleteById(id int) {
	log.Println("=== DELETE ITEM BY ID ===")

	_, err := itemRepo.db.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
