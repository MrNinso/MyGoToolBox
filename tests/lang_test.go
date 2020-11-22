package tests

import (
	"github.com/MrNinso/MyGoToolBox/lang/crypto"
	"github.com/MrNinso/MyGoToolBox/lang/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

// crypto tests
func TestCryptoMD5(t *testing.T) {
	assert.Equal(t, "6a35e5064b0084772b24e448f7d9d674", crypto.ToMD5Hash("wtf Man"))
}

func TestCryptoSHA256(t *testing.T) {
	assert.Equal(t,
		"d943277785598d1eb21e1ccfae040f4aec98f8e74db3fffa4c408d5934e5c25d",
		crypto.ToSHA256Hash("wtf Man"),
	)
}

// env test
func TestSetAnGetValue(t *testing.T) {
	_, err := env.SetEnv("TEST", "AAAAAAAAAAAAAAAAAAAAa")

	assert.Nil(t, err)

	h, err := env.SetEnv("TEST", "AAAAAAAAAAAAAAAAAAAAa")

	assert.Nil(t, err)
	assert.True(t, h)

	assert.Equal(t, "AAAAAAAAAAAAAAAAAAAAa", env.GetEnv("TEST", ""))
	assert.Equal(t, "B", env.GetEnv("TEST2", "B"))

}

func TestGetInt(t *testing.T) {
	e, err := env.SetEnv("TestGetInt1", "AaaaaaaaaaaaaaaaaaaaaaaaA")

	assert.Nil(t, err)
	assert.False(t, e)

	e, err = env.SetEnv("TestGetInt2", "123")

	assert.Nil(t, err)
	assert.False(t, e)

	assert.Equal(t, -1, env.GetIntFromEnv("TestGetInt1", -1))
	assert.Equal(t, 123, env.GetIntFromEnv("TestGetInt2", -1))
}
