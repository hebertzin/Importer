package routes

import (
	"enube-challenge/packages/controllers"
	"enube-challenge/packages/repository"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Suppliers(router *gin.Engine, db *gorm.DB) {
	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := services.NewSupplierService(supplierRepository)
	supplierController := controllers.NewSupplierController(supplierService)

	supplierGroup := router.Group("/api/v1")
	{
		supplierGroup.POST("/suppliers/import", supplierController.ImportSuppliersHandler)
		supplierGroup.GET("/suppliers", supplierController.FindSuppliersHandler)
		supplierGroup.GET("/suppliers/:id", supplierController.FindSupplierById)
	}
}
