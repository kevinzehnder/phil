package config

import (
	"fmt"
	"os"
	"testing"
)

func TestConfiguration(t *testing.T) {
	os.Setenv("Something", "override")
	settings, err := GetConfig()
	if err != nil {
		t.Errorf("configuration error: %s", err)
	}

	fmt.Println(settings)
}
