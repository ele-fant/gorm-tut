package main

import (
	"fmt"
	"log"

	"gorm-tut/database"
	"gorm-tut/models"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	ps := database.NewPostgresService(db)

	prod := &models.Product{}
	err = ps.GetProductByID(prod, 1)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(*prod)

	// get list of existing tables
	var tables []string
	if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	//fmt.Println(tables)
}
