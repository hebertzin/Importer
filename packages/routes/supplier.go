package routes

import (
	"enube-challenge/packages/controllers"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Importer(router *gin.Engine, db *gorm.DB) {
	supplierService := services.NewSupplierService(db)
	supplierController := controllers.NewSupplierController(supplierService)

	supplierGroup := router.Group("/api/v1")
	{
		supplierGroup.POST("/import/suppliers", supplierController.ImportSuppliersHandler)
	}
}
