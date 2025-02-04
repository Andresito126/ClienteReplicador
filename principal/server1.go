package principal

import (
	"sync"
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


