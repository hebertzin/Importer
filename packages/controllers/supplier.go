package controllers

import (
	"bytes"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
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

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}(file)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	log.Printf("File read successfully, size: %d bytes", buf.Len())

	if err := ctrl.service.ImportSuppliersFromFile(c.Request.Context(), buf.Bytes()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to import suppliers: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Suppliers imported successfully"})
}
