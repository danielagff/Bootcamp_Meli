package service

import (
	"aula2gobases/internal/model"
	"aula2gobases/internal/repository"
	"aula2gobases/pkg/utils"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	if err := product.Validate(); err != nil {
		return model.Product{}, err
	}

	product.CodeValue = getProductCodeValue(s.repo.DataBase.Products)

	return s.repo.Save(product)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductById(id int) (model.Product, error) {
	return s.repo.GetById(id)
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

func (s *ProductService) UpdateProductPriceById(id int, product model.Product) (model.Product, error) {

	_, err := s.repo.GetById(id)
	if err != nil {
		return model.Product{}, err
	}

	return s.repo.UpdatePriceById(id, product)
}

func  getProductCodeValue(products map[int]model.Product) string {
	existsCode := make(map[string]bool)
	for _, product := range  products {
		existsCode[product.CodeValue] = true
	}
	return utils.GenerateProductCodeValue(existsCode)
}
