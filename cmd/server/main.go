package main

import (
	"fmt"
	"github.com/verma-kartik/marketplace-api/internal/database"
	_ "github.com/verma-kartik/marketplace-api/internal/database"
	"github.com/verma-kartik/marketplace-api/internal/services"
	"github.com/verma-kartik/marketplace-api/transport/http"
	_ "github.com/verma-kartik/marketplace-api/transport/http"
)

func Run() error {
	fmt.Println("starting the application")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to db")
		return err
	}

	//db.MigrateDb()

	//productService := models.NewProductService(db)
	//fmt.Println(productService.GetProduct(
	//	context.Background(),
	//	1))
	//
	//prod := models.Product{
	//	Name:         "Samsung S23",
	//	SerialNumber: 2,
	//	Quantity:     21,
	//	Price:        499.99,
	//	Description:  "flagship samsung product",
	//}
	//productService.CreateProduct(
	//	context.Background(),
	//	&prod)

	productService := services.NewProductService(db)

	httpHandler := http.NewHandler(productService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("hi im inside main")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
