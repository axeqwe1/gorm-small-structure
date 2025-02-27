package service

import (
	"coffee-shop/dto/request"
	"coffee-shop/internal/model"
	"coffee-shop/internal/repository"
	"fmt"
	"strconv"
)

func GetProductList() ([]model.Product, error) {
	return repository.GetProduct()
}

func CreateProductService(req request.ProductRequest) (*model.Product, error) {
	product := model.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	return repository.CreateProduct(&product)
}
func UpdateProductService(ID string, req request.ProductRequest) (*model.Product, error) {
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		// ถ้ามีข้อผิดพลาดเกิดขึ้น
		fmt.Println("Error parsing ID:", err)
		return nil, nil
	}
	product := model.Product{
		ID:    uint(id),
		Name:  req.Name,
		Price: req.Price,
	}

	return repository.UpdateProduct(&product)
}

// DeleteProductService
func DeleteProductService(ID string) error {
	// แปลง ID จาก string เป็น uint
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing ID: %v", err)
	}

	// เรียกใช้ repository เพื่อลบข้อมูลจากฐานข้อมูล
	return repository.DeleteProduct(uint(id))
}
