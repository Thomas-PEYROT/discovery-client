package discovery_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type RegisterMicroserviceBody struct {
	Name string `json:"name"`
}

type RegisterMicroserviceResponse struct {
	UUID string `json:"uuid"`
	Port uint32 `json:"port"`
}

func RegisterMicroservice(name string) {
	godotenv.Load()
	fmt.Println("Attempt to register microservice at ", os.Getenv("DISCOVERY_SERVER_URL"))

	url := os.Getenv("DISCOVERY_SERVER_URL") + "/register"
	body := RegisterMicroserviceBody{Name: name}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("erreur lors de la conversion en JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("erreur lors de la requête POST: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("réponse du serveur avec un statut non-OK: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("erreur lors de la lecture du corps de la réponse: %v", err)
	}

	var response RegisterMicroserviceResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("erreur lors du décodage JSON: %v", err)
	}

	if err != nil {
		log.Fatalf("Error registering microservice: %v", err)
	}
	ServiceInformations = response

	fmt.Println("Microservice registered successfully", response)
}
