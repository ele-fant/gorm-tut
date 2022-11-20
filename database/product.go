package database

import (
	"gorm-tut/models"
	"log"

	"gorm.io/gorm"
)

func (g *PostgresService) GetProductByID(prod *models.Product, id int) error {

	err := g.DB.First(&prod, id)
	if err.Error == gorm.ErrRecordNotFound {
		log.Println("No record found in products for id", id)
		return err.Error
	}
	if err != nil {
		return err.Error
	}
	return nil
}
