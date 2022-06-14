package lib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/spf13/viper"
)

func TestLoadEnvironment(t *testing.T) {
	os.Setenv("ENVIRONMENT_SIMULATION", "1")
	data := []byte(`
SIMPLE_DATA=
SAMPLE_DATA=123
`)
	ioutil.WriteFile(".env", data, 0644)
	LoadEnvironment(map[string]interface{}{
		"simple_data": "123",
		"sample_data": "123",
	})
	defer os.Remove(".env")

	utils.AssertEqual(t, "123", viper.GetString("SAMPLE_DATA"))
	utils.AssertEqual(t, "123", os.Getenv("SAMPLE_DATA"))

}
