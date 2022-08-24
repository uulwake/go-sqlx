package main

import (
	"fmt"
	"go-sqlx/models"
	"go-sqlx/repositories"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "secret"
		dbname   = "go-sql"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	ItemRepository := repositories.NewItemRepository(db)
	count := ItemRepository.CountAll()
	fmt.Println(count)

	newItem := models.Item{Name: "itemZ", Qty: 10, Weight: 20}
	ItemRepository.Create(newItem)

	newItems := []models.Item{
		{Name: "itemA1", Qty: 1, Weight: 1},
		{Name: "itemA2", Qty: 2, Weight: 2},
		{Name: "itemA3", Qty: 3, Weight: 3},
		{Name: "itemA4", Qty: 4, Weight: 4},
		{Name: "itemA5", Qty: 5, Weight: 5},
	}
	ItemRepository.CreateBatch(newItems)

	items := ItemRepository.FetchAll()
	fmt.Println(items)

	updatedItem := items[0]
	updatedItem.Name = "newName"

	item := ItemRepository.FetchById(updatedItem.ID)
	fmt.Println(item)

	ItemRepository.UpdateById(updatedItem.ID, updatedItem)

	item = ItemRepository.FetchById(updatedItem.ID)
	fmt.Println(item)

	count = ItemRepository.CountAll()
	fmt.Println("Total items before delete:", count)

	lastId := items[len(items)-1].ID
	ItemRepository.DeleteById(lastId)

	count = ItemRepository.CountAll()
	fmt.Println("Total items after delete:", count)
}
