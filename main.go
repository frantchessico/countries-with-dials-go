package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dial_code"`
	Code     string `json:"code"`
}

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	r := gin.Default()

	// Define a porta padrão
	port := getPort()

	// Endpoint para retornar o arquivo JSON
	r.GET("/countries-with-dialing-code", func(c *gin.Context) {
		// Abra e leia o arquivo JSON
		file, err := os.Open("datas.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		var countries []Country
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&countries); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, countries)
	})

	// Inicie o servidor na porta definida
	r.Run(":" + port)
}

// getPort obtém a porta da variável de ambiente PORT ou usa uma porta padrão (8080)
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	return port
}
