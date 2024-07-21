package routes

import (
	"enube-challenge/packages/controllers"
	"enube-challenge/packages/repository"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Importer(router *gin.Engine, db *gorm.DB) {

	supplierRepository := repository.NewSupplierRepository(db)

	supplierService := services.NewSupplierService(supplierRepository)
	supplierController := controllers.NewSupplierController(supplierService)

	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/import/suppliers", supplierController.ImportSuppliersHandler)
	}
}
