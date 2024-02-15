package main

import (
	"github.com/DevHardiansyah/BasicGoRestAPI/controllers/productcontroller"
	"github.com/DevHardiansyah/BasicGoRestAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products/:id", productcontroller.Delete)

	r.Run()
}
