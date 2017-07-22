package uhandlers

import "encoding/json"

// ToJSON help to marshal object to json strings
func ToJSON(object interface{}) string {
	response, err := json.Marshal(object)
	if err != nil {
		return "{\"status\": \"failed\"}"
	}
	return string(response)
}
