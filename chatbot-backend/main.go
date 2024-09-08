package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// go run main.go para iniciar el file

type Response struct {
	Keywords []string `json:"keywords"`
	Response string   `json:"response"`
	Default  bool     `json:"default"`
}

// Cargar el JSON de respuestas [a fines de pruebas]
func loadResponses() ([]Response, error) {
	file, err := os.Open("responses.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var responses []Response
	if err := json.NewDecoder(file).Decode(&responses); err != nil {
		return nil, err
	}

	return responses, nil
}

// Función que busca la respuesta basada en la pregunta del usuario
func findResponse(question string, responses []Response) string {
	question = strings.ToLower(question)

	for _, response := range responses {
		for _, keyword := range response.Keywords {
			if strings.Contains(question, keyword) {
				return response.Response
			}
		}
	}

	// Si no se encuentra ninguna palabra clave, devolver una respuesta por defecto
	for _, response := range responses {
		if response.Default {
			return response.Response
		}
	}

	return "No tengo una respuesta para eso."
}

func main() {
	// 	r refers to the router
	// c refers to the request context
	r := gin.Default()

	// CORS middleware para permitir todas las solicitudes [a fines de prueba]
	r.Use(cors.Default())

	responses, err := loadResponses()
	if err != nil {
		fmt.Println("Error cargando respuestas:", err)
		return
	}

	// Endpoint para recibir las preguntas del usuario
	r.POST("/chatbot", func(c *gin.Context) {
		var jsonBody struct {
			Question string `json:"question"`
		}

		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Petición inválida"})
			return
		}

		answer := findResponse(jsonBody.Question, responses)

		// Responder en formato JSON
		c.JSON(http.StatusOK, gin.H{
			"response": answer,
		})

	})

	r.Run(":8080")
}
