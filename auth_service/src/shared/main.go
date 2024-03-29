package shared

import (
	"encoding/json"
)

// StructToMap converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return 
}
