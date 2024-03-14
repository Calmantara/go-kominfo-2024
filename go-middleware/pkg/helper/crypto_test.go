package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	t.Run("error genereate hash", func(t *testing.T) {
		dummyPass := make([]byte, 80)
		res, err := GenerateHash(string(dummyPass))
		// exptected ku err != nil
		assert.NotNil(t, err)
		// res == ""
		assert.Equal(t, "", res)
	})

	t.Run("success generate hash", func(t *testing.T) {
		dummyPass := make([]byte, 10)
		res, err := GenerateHash(string(dummyPass))
		// exptected ku err != nil
		assert.Nil(t, err)
		// res == ""
		assert.NotEqual(t, "", res)
	})
}
