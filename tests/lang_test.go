package tests

import (
	"github.com/MrNinso/MyGoToolBox/lang/crypto"
	"github.com/MrNinso/MyGoToolBox/lang/env"
	"github.com/MrNinso/MyGoToolBox/lang/ifs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	T *testing.T
}

// crypto tests
func (T *testStruct) TestCryptoMD5() {
	assert.Equal(T.T, "bbe9f2c0b66df3c89869a4be4313000f", crypto.ToMD5Hash("wtf Man"))
}

func (T *testStruct) TestCryptoSHA256() {
	assert.Equal(T.T,
		"32e8e4b2a16368d2750e33393ff897dd5d28cd74ba6fe8827e7c4e658a8f543a",
		crypto.ToSHA256Hash("wtf Man"),
	)
}

// env test
func (T *testStruct) SetAnGetValue() {
	_, err := env.SetEnv("TEST", "AAAAAAAAAAAAAAAAAAAAa")

	assert.Nil(T.T, err)

	h, err := env.SetEnv("TEST", "AAAAAAAAAAAAAAAAAAAAa")

	assert.Nil(T.T, err)
	assert.True(T.T, h)

	assert.Equal(T.T, "AAAAAAAAAAAAAAAAAAAAa", env.GetEnv("TEST", ""))
	assert.Equal(T.T, "B", env.GetEnv("TEST2", "B"))

}

// Ifs Tests
func (T *testStruct) IfReturn() {
	assert.True(T.T, ifs.IfReturn(true, true, false).(bool))
	assert.False(T.T, ifs.IfReturn(false, true, false).(bool))
}

func (T *testStruct) IfNil() {
	assert.Equal(T.T, "1", ifs.IfNil("1", "2").(string))
	assert.Equal(T.T, "2", ifs.IfNil(nil, "2").(string))
}

func (T testStruct) IfStringEmpty() {
	assert.Equal(T.T, "1", ifs.IfStringEmpty("1", "2"))
	assert.Equal(T.T, "2", ifs.IfStringEmpty("", "2"))
}

func (T *testStruct) IfAnyNil() {
	assert.True(T.T, ifs.IfAnyNil("", "", "", nil, ""))
	assert.False(T.T, ifs.IfAnyNil("", "", "", "", ""))
}

func (T *testStruct) IfAnyStringEmpty() {
	assert.True(T.T, ifs.IfAnyStringEmpty("1", "2", "", "4", "5"))
	assert.False(T.T, ifs.IfAnyStringEmpty("1", "2", "3", "4", "5"))
}

func (T *testStruct) IfStringArrayItemContains() {
	strs := []string{"aaa1aa", "bb", "C", "d"}

	t, p := ifs.IfStringArrayItemContains("1", strs)

	assert.True(T.T, t)
	assert.Equal(T.T, 0, p)

	t, p = ifs.IfStringArrayItemContains("bb", strs)

	assert.True(T.T, t)
	assert.Equal(T.T, 1, p)

	t, p = ifs.IfStringArrayItemContains("e", strs)

	assert.False(T.T, t)
	assert.Equal(T.T, -1, p)
}
