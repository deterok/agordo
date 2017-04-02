package loaders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = map[interface{}]interface{}{
	"field_1": map[interface{}]interface{}{
		"field_1_1": "value_1_1",
		"field_1_2": map[interface{}]interface{}{
			"field_1_2_1": 1,
			"field_1_2_2": "string",
			"field_1_2_3": []interface{}{
				1.,
				2,
				map[interface{}]interface{}{
					"2.0":     2.0,
					"2.2":     2.2,
					"2.4":     2.4,
					"2.6":     2.6,
					"2.8":     2.8,
					"2.99999": 2.99999,
				},
				3.,
				4,
				map[interface{}]interface{}{
					"4.0":     4.0,
					"4.2":     4.2,
					"4.4":     4.4,
					"4.6":     4.6,
					"4.8":     4.8,
					"4.99999": 4.99999,
				},
				[]interface{}{
					map[interface{}]interface{}{
						"slice_1": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_2": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_3": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_4": map[interface{}]interface{}{
							"1": 1,
							"2": 2,
							"3": 3,
							"4": 4,
							"5": 5,
							"6": 6,
							"7": 7,
							"8": 8,
							"9": 9,
							"0": 0,
						},
						"slice_5": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
					},
				},
			},
			"field_1_2_4": true,
			"field_1_2_5": 0.12,
		},
	},
	"field_2_1": map[interface{}]interface{}{
		"field_2_1_1": map[interface{}]interface{}{
			"field_2_1_1_1": map[interface{}]interface{}{},
			"slice":         []interface{}{"1", "2", "3", "4", "5", 1, 2, 3, 4, 5},
		},
	},
}

var convertedTestData = map[string]interface{}{
	"field_1": map[string]interface{}{
		"field_1_1": "value_1_1",
		"field_1_2": map[string]interface{}{
			"field_1_2_1": 1,
			"field_1_2_2": "string",
			"field_1_2_3": []interface{}{
				1.,
				2,
				map[string]interface{}{
					"2.0":     2.0,
					"2.2":     2.2,
					"2.4":     2.4,
					"2.6":     2.6,
					"2.8":     2.8,
					"2.99999": 2.99999,
				},
				3.,
				4,
				map[string]interface{}{
					"4.0":     4.0,
					"4.2":     4.2,
					"4.4":     4.4,
					"4.6":     4.6,
					"4.8":     4.8,
					"4.99999": 4.99999,
				},
				[]interface{}{
					map[string]interface{}{
						"slice_1": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_2": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_3": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
						"slice_4": map[string]interface{}{
							"1": 1,
							"2": 2,
							"3": 3,
							"4": 4,
							"5": 5,
							"6": 6,
							"7": 7,
							"8": 8,
							"9": 9,
							"0": 0,
						},
						"slice_5": []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
					},
				},
			},
			"field_1_2_4": true,
			"field_1_2_5": 0.12,
		},
	},
	"field_2_1": map[string]interface{}{
		"field_2_1_1": map[string]interface{}{
			"field_2_1_1_1": map[string]interface{}{},
			"slice":         []interface{}{"1", "2", "3", "4", "5", 1, 2, 3, 4, 5},
		},
	},
}

func TestConvertMap(t *testing.T) {
	procData, err := convertMap(testData)
	if assert.NoError(t, err) {
		assert.Equal(t, convertedTestData, procData)
	}
}
