package discovery_client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type UnregisterMicroserviceBody struct {
	UUID string `json:"uuid"`
}

type UnregisterMicroserviceResponse struct {
	Message string `json:"message"`
}

func UnregisterMicroservice() {
	url := os.Getenv("DISCOVERY_SERVER_URL") + "/unregister"
	body := UnregisterMicroserviceBody{UUID: ServiceInformations.UUID}

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

	var response UnregisterMicroserviceResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		log.Fatalf("erreur lors du décodage JSON: %v", err)
	}
}
