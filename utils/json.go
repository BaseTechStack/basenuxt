package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// ReplaceJSONField replaces a field in a JSON string with a new value
func ReplaceJSONField(jsonStr, field, newValue string) string {
	// Try to parse and modify using the json package first
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err == nil {
		jsonMap[field] = json.RawMessage(newValue)
		if newJSON, err := json.MarshalIndent(jsonMap, "", "  "); err == nil {
			return string(newJSON)
		}
	}

	// Fallback to regex replacement if JSON parsing fails
	pattern := fmt.Sprintf(`"?%s"?\s*:\s*"[^"]*"`, field)
	re := regexp.MustCompile(pattern)
	replacement := fmt.Sprintf(`"%s": %s`, field, newValue)
	return re.ReplaceAllString(jsonStr, replacement)
}
