package agordo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testJSONFixture = map[string]interface{}{
	"dev": map[string]interface{}{
		"parametr_1": []interface{}{"test_value_1", "test_value_2", "test_value_3", "test_value_4"},
		"parametr_2": map[string]interface{}{
			"subparam_1": 1.5091,
			"subparam_2": 0.0001,
			"subparam_3": "0.413"},
		"parametr_3": []interface{}{
			map[string]interface{}{"value": 1., "str": "1"},
			map[string]interface{}{"value": 2., "str": "2"},
			map[string]interface{}{"value": 3., "str": "3"},
			map[string]interface{}{"value": 4., "str": "4"},
			map[string]interface{}{"value": 5., "str": "5", "bool": true},
		},
		"parametr_4": true,
	},
}

func TestDefaultConfig(t *testing.T) {
	conf := New()
	conf.SetDefaults(testJSONFixture)
	conf.Assemble()

	t.Run("GetOnlyDefaults", func(t *testing.T) {
		confMap := conf.Defaults()
		assert.EqualValues(t, testJSONFixture, confMap)
	})

	t.Run("GetDefaultsFromAllConfig", func(t *testing.T) {
		confMap := conf.GetAll()
		assert.EqualValues(t, testJSONFixture, confMap)
	})

	t.Run("ClearDefaults", func(t *testing.T) {
		conf.ClearDefaults()
		conf.Assemble()

		confMap := conf.GetAll()
		assert.Equal(t, map[string]interface{}{}, confMap)
	})
}

func TestSeparator(t *testing.T) {
	conf := New()

	sepValue := "<$+$>"
	conf.SetSeparator(sepValue)

	assert.Equal(t, sepValue, conf.Separator())
}

func TestGetAll(t *testing.T) {
	conf := New()
	conf.AddPaths("tests/test_data.json")
	conf.Assemble()

	confMap := conf.GetAll()
	assert.Equal(t, testJSONFixture, confMap)
}

func TestGets(t *testing.T) {
	conf := New()
	conf.AddPaths("tests/test_data.json")
	conf.Assemble()

	t.Run("GetFirstLvl", func(t *testing.T) {
		value := conf.Get("dev")
		assert.Equal(t, testJSONFixture["dev"], value)
	})

	t.Run("GetWithCompositeKey", func(t *testing.T) {
		value := conf.Get("dev.parametr_2.subparam_3")
		assert.Equal(t, "0.413", value)
	})

	t.Run("GetWithNonExistentKey", func(t *testing.T) {
		value := conf.Get("dev.parametr_2.Non-ExistentKey")
		assert.Nil(t, value)
	})

	t.Run("GetBool", func(t *testing.T) {
		value := conf.GetBool("dev.parametr_4")

		assert.Equal(t, true, value)

		value = conf.GetBool("dev")
		assert.Equal(t, false, value)

		value = conf.GetBool("dev.parametr_2.Non-ExistentKey")
		assert.Equal(t, false, value)
	})

	t.Run("GetString", func(t *testing.T) {
		value := conf.GetString("dev.parametr_2.subparam_3")

		assert.Equal(t, "0.413", value)

		value = conf.GetString("dev")
		assert.Equal(t, "", value)

		value = conf.GetString("dev.parametr_2.Non-ExistentKey")
		assert.Equal(t, "", value)
	})

	t.Run("GetInt", func(t *testing.T) {
		value := conf.GetInt("dev.parametr_2.subparam_1")

		assert.Equal(t, 1, value)

		value = conf.GetInt("dev")
		assert.Equal(t, 0, value)

		value = conf.GetInt("dev.parametr_2.Non-ExistentKey")
		assert.Equal(t, 0, value)
	})
}
