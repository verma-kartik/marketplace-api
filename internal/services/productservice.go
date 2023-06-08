package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/verma-kartik/marketplace-api/internal/models"
)

var (
	ErrFetchingProduct = errors.New("failed to fetch the product by serial number")
	//ErrFetchingProducts = errors.New("failed to fetch the products")
)

type ProductService struct {
	_productRepository models.ProductRepository
}

// NewProductService ctor-composite literal
// return pointer to a NewProductService
func NewProductService(productRepository models.ProductRepository) *ProductService {
	return &ProductService{
		_productRepository: productRepository,
	}
}

func (p *ProductService) GetProduct(ctx context.Context, serialNumber int32) (models.Product, error) {
	fmt.Println("retrieving a product")
	product, err := p._productRepository.GetProductBySerialN(ctx, serialNumber)
	if err != nil {
		fmt.Println(err)
		return models.Product{}, ErrFetchingProduct
	}
	return product, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	fmt.Println("creating a new product")
	err := p._productRepository.CreateProduct(ctx, product)
	if err != nil {
		fmt.Println("could not add product")
		return err
	}
	fmt.Println("successfully created product")

	return nil
}
