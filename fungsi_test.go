package fungsi

import (
	"testing"
)

func TestExpandToMap(t *testing.T) {
	keys := "api.server.host"
	sep := "."
	value := "localhost"

	result := ExpandToMap(keys, sep, value, nil)

	if v, ok := result["api"].(map[string]interface{})["server"].(map[string]interface{})["host"]; !ok || v != "localhost" {
		t.Errorf("Test failed, returned map is not valid\n")
	}
}
func TestExpandToMapWithInit(t *testing.T) {
	keys := "api.server.host"
	sep := "."
	value := "localhost"
	init := map[string]interface{}{
		"api": map[string]interface{}{
			"server": map[string]interface{}{
				"port": 8080,
			},
		},
	}

	result := ExpandToMap(keys, sep, value, init)

	if v, ok := result["api"].(map[string]interface{})["server"].(map[string]interface{})["port"]; !ok || v != 8080 {
		t.Errorf("Test failed, returned map is not valid\n")
	}
}

func TestFlattenMap(t *testing.T) {
	m := map[string]interface{}{
		"api": map[string]interface{}{
			"server": map[string]interface{}{
				"host": "localhost",
				"port": 8080,
			},
		},
	}
	result := FlattenMap(m)

	if v, ok := result["api.server.host"]; !ok || v != "localhost" {
		t.Errorf("Test failed, returned map is not valid\n")
	}

	if v, ok := result["api.server.port"]; !ok || v != 8080 {
		t.Errorf("Test failed, returned map is not valid\n")
	}
}
