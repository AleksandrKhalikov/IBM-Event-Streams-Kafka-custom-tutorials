package helper_functions

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrettyPrintJSON prints a JSON message in a human-readable format.
func PrettyPrintJSON(message []byte) {
	var prettyJSON map[string]any
	if err := json.Unmarshal(message, &prettyJSON); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	prettyBytes, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	fmt.Println(string(prettyBytes))
}
