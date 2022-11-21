package helper_test

import (
	"shortlink/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashedPassword(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		passwod := "password"
		test, err := helper.HashPassword(passwod)
		assert.Empty(t, err)
		assert.IsType(t, test, passwod)
	})
}

func TestMatchPassword(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		passwod := "password"
		hashPassword, err := helper.HashPassword(passwod)
		assert.Empty(t, err)
		assert.IsType(t, hashPassword, passwod)

		matchPassword := helper.MatchPassword(hashPassword, passwod)
		assert.Equal(t, matchPassword, true)
	})
}
