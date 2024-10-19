package utils

import (
	"encoding/json"
	"log"
)

const MESASGE = "Hello from server"

func WsSuccessResponse(data []byte) ([]byte, error) {
	unMarData := string(data)

	res := map[string]interface{}{
		"message":      MESASGE,
		"dataRecieved": unMarData,
	}

	resStr, error := json.Marshal(res)
	if error != nil {
		log.Println("Error marshaling JSON:", error)
		return nil, error
	}

	return resStr, nil
}

func WsErrorResponse() ([]byte, error) {
	res := map[string]interface{}{
		"message": "Error processing request",
	}

	resStr, error := json.Marshal(res)
	if error != nil {
		log.Println("Error marshaling JSON:", error)
		return nil, error
	}

	return resStr, nil
}
