package routes

import (
	"enube-challenge/packages/controllers"
	middleware "enube-challenge/packages/middlewares"
	"enube-challenge/packages/repository"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Suppliers(router *gin.Engine, db *gorm.DB) {

	jwtService := services.NewJWTService()

	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := services.NewSupplierService(supplierRepository)
	supplierController := controllers.NewSupplierController(supplierService)

	supplierGroup := router.Group("/api/v1")
	{
		supplierGroup.POST("/suppliers/import", middleware.AuthMiddleware(jwtService), supplierController.ImportSuppliersHandler)
		supplierGroup.GET("/suppliers", middleware.AuthMiddleware(jwtService), supplierController.FindSuppliersHandler)
		supplierGroup.GET("/suppliers/:id", middleware.AuthMiddleware(jwtService), supplierController.FindSupplierById)
	}
}
