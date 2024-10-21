package routes

import (
	"enube-challenge/packages/infra/db/repository"
	"enube-challenge/packages/presentation/controllers"
	"enube-challenge/packages/presentation/middlewares"
	"enube-challenge/packages/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Suppliers(router *gin.Engine, db *gorm.DB) {
	jwtUseCase := usecases.NewJWTUseCase()
	supplierRepository := repository.NewSupplierRepository(db)
	supplierUseCase := usecases.NewSupplierUseCase(supplierRepository)
	supplierController := controllers.NewSupplierController(supplierUseCase)
	supplierGroup := router.Group("/api/v1")
	{
		supplierGroup.POST("/suppliers/import", middleware.AuthMiddleware(jwtUseCase), middleware.VerifyXLSXMiddleware, supplierController.ImportSuppliersHandler)
		supplierGroup.GET("/suppliers", middleware.AuthMiddleware(jwtUseCase), supplierController.FindSuppliersHandler)
		supplierGroup.GET("/suppliers/:id", middleware.AuthMiddleware(jwtUseCase), supplierController.FindSupplierById)
	}
}
