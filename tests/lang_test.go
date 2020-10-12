package tests

import (
	"github.com/MrNinso/MyGoToolBox/lang/crypto"
	"github.com/MrNinso/MyGoToolBox/lang/env"
	"github.com/MrNinso/MyGoToolBox/lang/ifs"
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

// Ifs Tests
func TestIfReturn(t *testing.T) {
	assert.True(t, ifs.IfReturn(true, true, false).(bool))
	assert.False(t, ifs.IfReturn(false, true, false).(bool))
}

func TestIfNil(t *testing.T) {
	assert.Equal(t, "1", ifs.IfNil("1", "2").(string))
	assert.Equal(t, "2", ifs.IfNil(nil, "2").(string))
}

func TestIfStringEmpty(t *testing.T) {
	assert.Equal(t, "1", ifs.IfStringEmpty("1", "2"))
	assert.Equal(t, "2", ifs.IfStringEmpty("", "2"))
}

func TestIfAnyNil(t *testing.T) {
	assert.True(t, ifs.IfAnyNil("", "", "", nil, ""))
	assert.False(t, ifs.IfAnyNil("", "", "", "", ""))
}

func TestIfAnyStringEmpty(t *testing.T) {
	assert.True(t, ifs.IfAnyStringEmpty("1", "2", "", "4", "5"))
	assert.False(t, ifs.IfAnyStringEmpty("1", "2", "3", "4", "5"))
}

func TestIfStringArrayItemContains(t *testing.T) {
	strs := []string{"aaa1aa", "bb", "C", "d"}

	t1, p := ifs.IfStringArrayItemContains("1", strs)

	assert.True(t, t1)
	assert.Equal(t, 0, p)

	t1, p = ifs.IfStringArrayItemContains("bb", strs)

	assert.True(t, t1)
	assert.Equal(t, 1, p)

	t1, p = ifs.IfStringArrayItemContains("e", strs)

	assert.False(t, t1)
	assert.Equal(t, -1, p)
}
