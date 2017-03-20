package agordo

// TODO: move this logic in a separate project

import (
	"fmt"
	"strings"
)

// KeyExistError is returned when a value for a requested key is missing.
type KeyExistError struct {
	Key string
}

func (e KeyExistError) Error() string {
	return fmt.Sprintf("key '%s' does not exist", e.Key)
}

// ValueNotMapError can be returned when the composite key is not compiled
// correctly and in an intermediate element of the key the value has a type
// different from map[string]interface{}.
type ValueNotMapError struct {
	Key string
}

func (e ValueNotMapError) Error() string {
	return fmt.Sprintf("A value obtained by '%s' key "+
		"is not 'map[string]interface{}'", e.Key)
}

// DeepGet allows get a valuef from a deeply located map by a composite key.
//
// Returns the corresponding error when the wrong key (KeyExistError) or
// missing value (ValueNotMapError).
func DeepGet(m map[string]interface{}, key, sep string) (interface{}, error) {
	splittedKeys := strings.Split(key, sep)
	return deepGet(m, splittedKeys)
}

func deepGet(m map[string]interface{}, keys []string) (interface{}, error) {
	keysCount := len(keys)

	if keysCount > 0 {

		value, ok := m[keys[0]]
		if value == nil || keysCount == 1 {
			if ok {
				return value, nil
			}

			return nil, KeyExistError{keys[0]}
		}

		if v, ok := value.(map[string]interface{}); ok {
			return deepGet(v, keys[1:])
		}

		return nil, ValueNotMapError{keys[0]}
	}

	return m, nil
}
