package principal

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// modelo
type Product struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Cantidad    int    `json:"cantidad"`
	CodigoBarras string `json:"codigo_barras"`
}

//como la bd
var products []Product
var lastID = 0
var hasChanges = false
var mu sync.Mutex 


func SetupRouter() *gin.Engine {
	r := gin.Default()

	//get productos
	r.GET("/productos", func(c *gin.Context) {
		mu.Lock()
		hasChanges = true 
		mu.Unlock()
		
		c.JSON(http.StatusOK, products)
	})
	
	//post ptoducts
	r.POST("/productos", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		lastID++
		newProduct.ID = lastID
		products = append(products, newProduct)
		hasChanges = true 
		mu.Unlock()

		c.JSON(http.StatusCreated, newProduct)
	})

	//obtener cambios
	r.GET("/estado", func(c *gin.Context) {
		mu.Lock()
		if hasChanges {
			hasChanges = false 
			mu.Unlock()
			c.JSON(http.StatusOK, gin.H{"estado": "cambio"})
		} else {
			mu.Unlock()
			c.JSON(http.StatusOK, gin.H{"estado": "sin_cambio"})
		}
	})

	return r

}