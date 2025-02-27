package controller

import (
	"coffee-shop/dto/request"
	"coffee-shop/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func GetProductsAllController(c *gin.Context) {
	products, err := service.GetProductList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body request.ProductRequest true "Create product"
// @Success 201 {object} model.Product
// @Router /products [post]
func CreateProductController(c *gin.Context) {
	var product request.ProductRequest

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//call service to create product
	createProduct, err := service.CreateProductService(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, createProduct)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"  // เพิ่มการระบุ id ใน path
// @Param product body request.ProductRequest true "Product update request"
// @Success 200 {object} model.Product
// @Router /products/{id} [put]
func UpdateProductController(c *gin.Context) {
	var product request.ProductRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updateProduct, err := service.UpdateProductService(id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to Update Product"})
		return
	}

	c.JSON(http.StatusOK, updateProduct)
}

// UpdateProduct godoc
// @Summary DELETE an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"  // เพิ่มการระบุ id ใน path
// @Success 200 {object} model.Product
// @Router /products/{id} [delete]
func DeleteProductController(c *gin.Context) {
	// รับค่า id จาก path parameter
	id := c.Param("id")

	// เรียกใช้ service เพื่อลบข้อมูล
	err := service.DeleteProductService(id)
	if err != nil {
		// หากเกิดข้อผิดพลาดในการลบข้อมูล
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่ง response ว่าลบสำเร็จ
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
