package repository

import (
	"coffee-shop/db"
	"coffee-shop/internal/model"
	"fmt"
)

func GetProduct() ([]model.Product, error) {
	var products []model.Product
	result := db.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *model.Product) (*model.Product, error) {
	result := db.DB.Create(product)
	return product, result.Error
}

func UpdateProduct(product *model.Product) (*model.Product, error) {
	var existingProduct model.Product // ใช้ struct ธรรมดา

	if err := db.DB.First(&existingProduct, product.ID).Error; err != nil {
		return nil, err
	}

	// อัปเดตเฉพาะฟิลด์ที่จำเป็น
	updates := map[string]interface{}{
		"name":  product.Name,
		"price": product.Price,
	}

	result := db.DB.Model(&existingProduct).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingProduct, nil // Return pointer ของ struct ที่อัปเดตแล้ว
}

// DeleteProduct (Repository)
func DeleteProduct(ID uint) error {
	var existingProduct model.Product

	// ตรวจสอบว่ามีข้อมูลหรือไม่
	if err := db.DB.First(&existingProduct, ID).Error; err != nil {
		return fmt.Errorf("Error finding product: %v", err)
	}

	// ลบข้อมูล
	result := db.DB.Delete(&existingProduct)
	if result.Error != nil {
		return fmt.Errorf("Error deleting product: %v", result.Error)
	}

	return nil
}
