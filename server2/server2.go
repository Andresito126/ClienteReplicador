package replicacion

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var estadoActual = "No hay cambios en principal"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/estado", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"estado": estadoActual})
	})

	return r
}

func InicioReplica() {
	for {
		time.Sleep(5 * time.Second)
		resp, err := http.Get("http://localhost:8080/estado")
		if err != nil {
			fmt.Println(" Error conectando con el servidor principal:", err)
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		var data map[string]string
		json.Unmarshal(body, &data)

		//ver estado del server principal
		if data["estado"] == "cambio" {
			estadoActual = " SI hay cambios en principal"
		} else {
			estadoActual = " No hay cambios en principal"
		}

		fmt.Println(estadoActual)
		resp.Body.Close()
	}
}