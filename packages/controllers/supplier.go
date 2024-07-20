package controllers

import (
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type SupplierController struct {
	service services.SupplierService
}

func NewSupplierController(service services.SupplierService) *SupplierController {
	return &SupplierController{service}
}

func (ctrl *SupplierController) ImportSuppliersHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file contents"})
		return
	}

	if err := ctrl.service.ImportSuppliersFromFile(c.Request.Context(), fileBytes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to import suppliers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Suppliers imported successfully"})
}
