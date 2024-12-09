package service

import (
	"aula2gobases/internal/model"
	iservice "aula2gobases/internal/service/interfaces"
	"aula2gobases/internal/repository/interfaces"
	"aula2gobases/pkg/utils"
)

type ProductService struct {
	repo interfaces.IProductsRepository
}

func NewProductService(repo interfaces.IProductsRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	if err := product.Validate(); err != nil {
		return model.Product{}, err
	}

	products, _ := s.repo.GetAll()
	product.CodeValue = getProductCodeValue(products)

	return s.repo.Save(product)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductById(id int) (model.Product, error) {
	return s.repo.GetById(id)
}

func (s *ProductService) GetHigherPriceProductsByPrice(price float64) ([]model.Product, error) {
	products, err := s.repo.GetAll() 
	if err != nil {
		return []model.Product{}, err
	}
	var filteredProducts []model.Product
	for _, product := range products {
		if product.Price > price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

func (s *ProductService) DeleteProductById(id int) (string, error) {
	return s.repo.DeleteById(id)
}

func (s *ProductService) UpdateProductById(id int, product model.Product) (string, error) {
	if err := product.Validate(); err != nil {
		return "", err
	}
	return s.repo.UpdateById(id, product)
}

func (s *ProductService) UpdatePriceById(id int, product model.Product) (model.Product, error) {

	_, err := s.repo.GetById(id)
	if err != nil {
		return model.Product{}, err
	}

	return s.repo.UpdatePriceById(id, product)
}

func  getProductCodeValue(products []model.Product) string {
	existsCode := make(map[string]bool)
	for _, product := range  products {
		existsCode[product.CodeValue] = true
	}
	return utils.GenerateProductCodeValue(existsCode)
}

var _ iservice.IProductsService = (*ProductService)(nil)
