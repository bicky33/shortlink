package config_test

import (
	"shortlink/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigApp(t *testing.T) {
	config := config.LoadConfigApp("../")
	t.Run("LoadConfigApp", func(t *testing.T) {
		assert.Equal(t, config.AppPort, 3000)
		assert.NotEmpty(t, config.AppPort)
	})
}
