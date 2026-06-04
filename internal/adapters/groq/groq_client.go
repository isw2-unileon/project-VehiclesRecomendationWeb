package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
)

type GroqClient struct {
	apiKey string
	client *http.Client
}

func NewGroqClient() *GroqClient {
	return &GroqClient{
		apiKey: os.Getenv("GROQ_API_KEY"),
		client: &http.Client{},
	}
}

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model    string        `json:"model"`
	Messages []groqMessage `json:"messages"`
}

type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (g *GroqClient) GenerateRecommendation(userPreferences string, availableCars []domain.Car) (string, error) {
	if g.apiKey == "" {
		return "", fmt.Errorf("GROQ_API_KEY environment variable is not set")
	}

	carsContext := ""
	for _, car := range availableCars {
		carsContext += fmt.Sprintf("- ID: %d, %s %s, Fuel: %s, Price: %.2f, HP: %d, Seats: %d\n",
			car.ID, car.Brand, car.Model, car.FuelType, car.Price, car.HorsePower, car.Seats)
	}

	systemPrompt := "You are a helpful car recommendation assistant. " +
		"Analyze the following request along with its custom rules, and use ONLY the provided car catalog to make your decision.\n\n" +
		"Available Catalog:\n" + carsContext

	reqBody := groqRequest{
		Model: "llama-3.1-8b-instant",
		Messages: []groqMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf("My preferences are: %s", userPreferences)},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+g.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// AQUÍ ESTÁ LA MAGIA: Ahora leerá la queja exacta de Groq
		errorBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("groq api error (status %d): %s", resp.StatusCode, string(errorBody))
	}

	var groqResp groqResponse
	if err := json.NewDecoder(resp.Body).Decode(&groqResp); err != nil {
		return "", err
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("no recommendations received from Groq")
	}

	return groqResp.Choices[0].Message.Content, nil
}
