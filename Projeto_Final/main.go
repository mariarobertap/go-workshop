package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func main() {
	r := gin.Default()

	// Rota para listar todos os produtos
	r.GET("/produto", getProducts)

	// Rota para criar um novo produto
	r.POST("/produto", saveProducts)

	// Rota para deletar um produto
	r.DELETE("/produto/:id", deleteProduct)

	r.Run(":8080")
}
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func saveProducts(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, gin.H{"product": newProduct})
}

func deleteProduct(c *gin.Context) {

	// Obtém o ID do produto da URL
	productId := c.Param("id")

	// Converte o ID para inteiro
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID Invalid"})
		return
	}

	// Encontra o índice do produto na lista
	index := -1
	for i, p := range products {
		if p.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto nao encontrado"})
		return
	}

	// Remove o produto da lista
	products = append(products[:index], products[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
