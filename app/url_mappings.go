package app

import (
	"github.com/DeVasu/tortoise/controllers/ping"
	"github.com/DeVasu/tortoise/controllers/plan"
	// "github.com/DeVasu/tortoise/controllers/products"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("admin/plan", plan.Create)
	router.GET("admin/plan", plan.List)
	router.PUT("admin/plan/:planId/addPromo", plan.AddPromotion)
	router.DELETE("admin/plan/:planId/deletePromo", plan.DeletePromotion)

	router.GET("/plan", plan.List)
	// router.POST("/products", products.Create)
	// router.GET("/products", products.List)
	// router.GET("/products/:productId", products.GetById)
	// router.PUT("/products/:productId", products.Update)
	// router.DELETE("/products/:productId", products.Delete)

}
