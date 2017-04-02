package loaders

import "fmt"

type KeyNotStringError struct {
	Key interface{}
}

func (e KeyNotStringError) Error() string {
	return fmt.Sprintf("key %v is not string", e.Key)
}

func convertMap(input map[interface{}]interface{}) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	var convValue interface{}
	var err error

	for key, value := range input {
		if strKey, ok := key.(string); ok {

			switch v := value.(type) {
			default:
				convValue = v

			case map[interface{}]interface{}:
				convValue, err = convertMap(v)
				if err != nil {
					return nil, err
				}

			case []interface{}:
				convValue, err = convertSlice(v)
				if err != nil {
					return nil, err
				}
			}

			result[strKey] = convValue

		} else {
			return nil, KeyNotStringError{key}
		}
	}

	return result, nil
}

func convertSlice(input []interface{}) ([]interface{}, error) {
	var err error

	for i, value := range input {

		switch v := value.(type) {

		case []interface{}:
			input[i], err = convertSlice(v)
			if err != nil {
				return nil, err
			}

		case map[interface{}]interface{}:
			input[i], err = convertMap(v)
			if err != nil {
				return nil, err
			}

		}
	}

	return input, nil
}
