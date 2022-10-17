package shared

import (
	"encoding/json"
	"fmt"
)

// StructToMap converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (map[string]interface{}) {
	fmt.Println(obj)
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	var newMap map[string]interface{}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	if err != nil {
		panic(err)
	}
	fmt.Println(newMap)
	return newMap
}
