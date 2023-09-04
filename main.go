package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)



type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dial_code"`
	Code     string `json:"code"`
}

func main() {
	r := gin.Default()

	// Endpoint para retornar o arquivo JSON
	r.GET("/infos", func(c *gin.Context) {
		// Abra e leia o arquivo JSON
		file, err := os.Open("datas.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		var datas []Country
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&datas); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, datas)
	})

	r.Run(":8080")
}
