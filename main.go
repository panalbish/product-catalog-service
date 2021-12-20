package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/products", addProduct)
	r.GET("/api/products/:id", getProduct)
	r.GET("/api/products", getAllProducts)
	return r
}

var products = make(map[string]Product)

func addProduct(ctx *gin.Context) {
	var newProduct Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		return
	}

	newProductID := uuid.New().String()
	products[newProductID] = newProduct
	ctx.IndentedJSON(http.StatusCreated, gin.H{newProductID: newProduct})
}

func getProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	if products[productID] == (Product{}) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, products[productID])
}

func getAllProducts(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, products)
}
