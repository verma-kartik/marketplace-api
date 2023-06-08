package database

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/verma-kartik/marketplace-api/internal/models"
	_ "github.com/verma-kartik/marketplace-api/internal/models"
)

func (d *Database) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product

	query := d.gClient.Select("products.*").Group("products.serial_number")

	if err := query.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (d *Database) GetProductBySerialN(ctx context.Context, serialNumber int32) (models.Product, error) {

	p := models.Product{}

	query := d.gClient.Select("products.*")
	err := query.Where(
		"products.serial_number = ?", serialNumber).First(&p).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		fmt.Println("product exists but could not retrieve it")
		return p, err
	}

	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("product does not exist")
		return p, nil
	}

	return p, nil
}

func (d *Database) DeleteProduct(ctx context.Context, serialNumber int32) error {
	var p models.Product

	if err := d.gClient.Where("serial_number = ?", serialNumber).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateProduct(ctx context.Context, p *models.Product) error {
	if err := d.gClient.Save(&p).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateProduct(ctx context.Context, p *models.Product) error {
	result := d.gClient.Create(&p)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
