package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	t.Run("error username", func(t *testing.T) {
		user := UserSignUp{Username: ""}
		err := user.Validate()

		assert.NotNil(t, err)
	})
}
