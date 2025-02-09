package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/interfaces/http/response"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	if err := h.productService.CreateProduct(&product); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.ListProducts()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get products", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Products retrieved successfully", products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := h.productService.GetProduct(uint(productID))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product retrieved successfully", product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}
	product.ID = uint(productID)

	if err := h.productService.UpdateProduct(&product); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	if err := h.productService.DeleteProduct(uint(productID)); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product deleted successfully", nil)
}
