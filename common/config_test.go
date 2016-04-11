package common

import (
	"os"
	"testing"

	"github.com/kelseyhightower/envconfig"
)

func patchEnv() {
	os.Setenv("API_BACKEND_DB_NAME", "test")
	os.Setenv("API_BACKEND_PASSWORD", "test")
	os.Setenv("API_BACKEND_URL", "test")
	os.Setenv("API_BACKEND_USERNAME", "test")
	os.Setenv("API_MODE", "test")
}

func TestGetConfigReturnsIdenticalObjectEachCall(t *testing.T) {
	patchEnv()

	config0 := GetConfig()
	config1 := GetConfig()
	config2 := &Config{}
	envconfig.Process("api", config2)

	if config0 != config1 {
		t.Log("Objects should be identical.")
		t.Fail()
	}
	if config1 == config2 {
		t.Log("Objects should not be identical.")
		t.Fail()
	}
}
