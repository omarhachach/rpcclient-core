package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanTxOutSetObject_ToJSON(t *testing.T) {
	obj := &ScanTxOutSetObject{
		Descriptor: "Hello",
		Desc:       "Hi",
		Range:      10,
		RangeN:     nil,
	}
	expected := "Hello"
	actual, err := obj.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))

	obj.Descriptor = ""
	expected = `{"desc":"Hi","range":10}`
	actual, err = obj.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))

	obj.Range = 0
	expected = `{"desc":"Hi"}`
	actual, err = obj.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))

	obj.RangeN = []int{10, 1000}
	expected = `{"desc":"Hi","range":[10,1000]}`
	actual, err = obj.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))

	obj.Range = 10
	actual, err = obj.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))

}

func TestScriptPubKey_GetAddress(t *testing.T) {
	obj := &ScriptPubKey{
		Address:   "Hi",
		Addresses: []string{"bye", "hello"},
	}
	assert.Equal(t, "Hi", obj.GetAddress())
	obj.Address = ""
	assert.Equal(t, "bye", obj.GetAddress())
}
