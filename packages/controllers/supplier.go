package controllers

import (
	"bytes"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/services"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierController struct {
	service services.SupplierService
}

func NewSupplierController(service services.SupplierService) *SupplierController {
	return &SupplierController{service}
}

// ImportSuppliersHandler godoc
// @Summary Import suppliers from an Excel file
// @Description Import suppliers from an Excel file. The file should be an Excel file (.xlsx) containing supplier data.
// @Tags suppliers
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Suppliers Excel file"
// @Success 200 {object} domain.HttpResponse "Suppliers imported successfully"
// @Failure 400 {object} domain.HttpResponse "Failed to read file"
// @Failure 500 {object} domain.HttpResponse "Failed to import suppliers"
// @Router /suppliers/import [post]
func (ctrl *SupplierController) ImportSuppliersHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to retrieve file from request: " + err.Error(),
		})
		return
	}
	defer func(file multipart.File) {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}(file)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusInternalServerError, domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to read file content: " + err.Error(),
		})
		return
	}

	log.Printf("File read successfully, size: %d bytes", buf.Len())

	if err := ctrl.service.ImportSuppliersFromFile(c.Request.Context(), buf.Bytes()); err != nil {
		log.Printf("ImportSuppliersFromFile error: %v", err)
		c.JSON(http.StatusInternalServerError, domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to import suppliers: " + err.Error(),
		})
		return
	}

	response := domain.HttpResponse{
		Message: "Suppliers imported successfully",
		Code:    http.StatusOK,
	}
	c.JSON(http.StatusOK, response)
}

// FindSuppliersHandler godoc
// @Summary Retrieve a list of suppliers with pagination
// @Description Get a paginated list of suppliers from the database
// @Tags suppliers
// @Accept  json
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of suppliers per page" default(10)
// @Success 200 {string} Supplier successfully found
// @Failure 400  {string} Invalid page number
// @Failure 500 {string} Failed to retrieve suppliers
// @Router /api/v1/suppliers/import [get]
func (ctrl *SupplierController) FindSuppliersHandler(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number"})
		return
	}

	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page size"})
		return
	}

	suppliers, err := ctrl.service.GetSuppliers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to retrieve suppliers: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":      page,
		"total":     len(suppliers),
		"code":      http.StatusOK,
		"suppliers": suppliers,
	})
}

// FindSupplierById godoc
// @Summary Retrieve a supplier
// @Description Get a supplier
// @Tags suppliers
// @Accept  json
// @Produce  json
// @Param id path int true "Supplier ID"
// @Success 200 {string} Supplier successfully found
// @Failure 500 {string} Failed to retrieve suppliers
// @Router /api/v1/suppliers/:id [get]
func (ctrl *SupplierController) FindSupplierById(ctx *gin.Context) {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	supplier, err := ctrl.service.FindSupplierById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Supplier not found"})
		return
	}

	response := domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Supplier successfully found",
		Body:    supplier,
	}
	ctx.JSON(http.StatusOK, response)
}
