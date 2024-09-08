package main

import (
	"testing"
)

// datos dummy para el test
func mockResponses() []Response {
	return []Response{
		{
			Keywords: []string{"hola", "buenas", "saludos"},
			Response: "¡Hola! ¿En qué puedo ayudarte hoy?",
			Default:  false,
		},
		{
			Keywords: []string{"adios", "hasta luego"},
			Response: "¡Adiós! ¡Que tengas un buen día!",
			Default:  false,
		},
		{
			Keywords: []string{},
			Response: "Lo siento, no entiendo tu pregunta.",
			Default:  true,
		},
	}
}

// Verifica que cuando la pregunta contiene una palabra clave (por ejemplo, "hola"), se devuelve la respuesta correspondiente
func TestFindResponseWithKeywords(t *testing.T) {
	responses := mockResponses()

	// Pregunta con palabra clave "hola"
	question := "hola, ¿cómo estás?"
	expected := "¡Hola! ¿En qué puedo ayudarte hoy?"
	got := findResponse(question, responses)

	if got != expected {
		t.Errorf("findResponse() = %v; want %v", got, expected)
	}
}

// Verifica que si la pregunta no contiene ninguna palabra clave reconocida, la respuesta por defecto es la que se devuelve
func TestFindResponseDefault(t *testing.T) {
	responses := mockResponses()

	// Pregunta sin coincidencia de palabras clave
	question := "¿Qué puedes hacer?"
	expected := "Lo siento, no entiendo tu pregunta."
	got := findResponse(question, responses)

	if got != expected {
		t.Errorf("findResponse() = %v; want %v", got, expected)
	}
}

// Verifica que la función no sea sensible a mayúsculas/minúsculas al buscar palabras clave.
func TestFindResponseCaseInsensitive(t *testing.T) {
	responses := mockResponses()

	// Pregunta en mayúsculas con palabra clave "HOLA"
	question := "HOLA, ¿QUÉ TAL?"
	expected := "¡Hola! ¿En qué puedo ayudarte hoy?"
	got := findResponse(question, responses)

	if got != expected {
		t.Errorf("findResponse() = %v; want %v", got, expected)
	}
}
