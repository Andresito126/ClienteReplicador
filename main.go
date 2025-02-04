package main

import (
	"log"

	"time"

	"github.com/Andresito126/go-servidor-replicador/principal"
	replicacion "github.com/Andresito126/go-servidor-replicador/server2"
)

func main() {
	// server main
	go func() {
		log.Println("Servidor principal corriendo en :8080")
		if err := principal.SetupRouter().Run(":8080"); err != nil {
			log.Fatal("Error en servidor principal:", err)
		}
	}()

	time.Sleep(2 * time.Second)

	
	//server de replica
	go func() {
		log.Println("🔄 Servidor de replica corriendo en :8081")
		if err := replicacion.SetupRouter().Run(":8081"); err != nil {
			log.Fatal("Error en servidor de replica:", err)
		}
	}()

	
	go replicacion.InicioReplica()

	
	select {}
}
