package agordo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMap = map[string]interface{}{
	"key_1": map[string]interface{}{
		"key_1_1": map[string]interface{}{
			"key_1_1_1": "value_1_1_1",
			"key_1_1_2": "value_1_1_2",
			"key_1_1_3": "value_1_1_3",
		},
		"key_1_2": map[string]interface{}{
			"key_1_2_1": "value_1_2_1",
			"key_1_2_2": "value_1_2_2",
			"key_1_2_3": "value_1_2_3",
		},
		"key_1_3": map[string]interface{}{
			"key_1_3_1": "value_1_3_1",
			"key_1_3_2": "value_1_3_2",
			"key_1_3_3": "value_1_3_3",
		},
	},
	"key_2": map[string]interface{}{
		"key_2_1": map[string]interface{}{
			"key_2_1_1": 211,
			"key_2_1_2": 212,
			"key_2_1_3": 213,
		},
		"key_2_2": map[string]interface{}{
			"key_2_2_1": 221,
			"key_2_2_2": 222,
			"key_2_2_3": 223,
		},
		"key_2_3": map[string]interface{}{
			"key_2_3_1": 231,
			"key_2_3_2": 232,
			"key_2_3_3": 233,
		},
	},
	"key_3": struct{ Test string }{
		Test: "test",
	},
}

const sep = "."

func TestDeepGet(t *testing.T) {
	key := "key_2.key_2_2.key_2_2_3"
	value, _ := DeepGet(testMap, key, sep)
	assert.Equal(t, 223, value.(int))
}

func TestDeepGetByNonExistentKey(t *testing.T) {
	key := "key_2.key_2_2.Non-ExistentKey"
	value, err := DeepGet(testMap, key, sep)
	assert.Equal(t, KeyExistError{"Non-ExistentKey"}, err)
	assert.Nil(t, value)
}

func TestDeepGetEmptyKey(t *testing.T) {
	value, err := DeepGet(testMap, "", sep)
	assert.Equal(t, KeyExistError{""}, err)
	assert.Nil(t, value)
}

func TestDeepGetFromNotMap(t *testing.T) {
	key := "key_2.key_2_2.key_2_2_3.Non-ExistentKey"
	value, err := DeepGet(testMap, key, sep)
	assert.Equal(t, ValueNotMapError{"key_2_2_3"}, err)
	assert.Nil(t, value)
}

func TestDeepGetComplicatedSeporator(t *testing.T) {
	sep := "$愛"
	key := "key_1$愛key_1_3$愛key_1_3_2"
	value, _ := DeepGet(testMap, key, sep)
	assert.Equal(t, "value_1_3_2", value)
}
