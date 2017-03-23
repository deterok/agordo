// Package agordo is a configuration system for your an application

package agordo

import (
	"github.com/deterok/agordo/loaders"

	"github.com/linkosmos/mapop"
	"github.com/mitchellh/mapstructure"
)

var config = new(Config)

// Config allows control a configurations of application.
type Config struct {
	conf      map[string]interface{}
	preset    map[string]interface{}
	loaders   []loaders.Loader
	separator string
}

// GetConfig return a default Config object.
func GetConfig() *Config {
	return config
}

// AddLoader connect a loader to a Config object.
func (c *Config) AddLoader(loader loaders.Loader) {
	c.loaders = append(c.loaders, loader)
}

// AddPaths is a "sugar" for adding a new file loader.
func (c *Config) AddPaths(paths ...string) {
	fl := loaders.NewFileLoader(paths...)
	c.AddLoader(fl)
	err := fl.Load()

	if err != nil {
		// TODO: Add error
	}
}

func (c Config) Defaults() map[string]interface{} {
	return c.preset
}

func (c *Config) SetDefaults(values map[string]interface{}) {
	c.preset = values
}

func (c Config) IsDefault(key string) bool {
	value := c.Get(key)
	defaultValue, _ := DeepGet(c.preset, key, c.separator)

	if value == defaultValue {
		return true
	}

	return false
}

func (c Config) ClearDefaults() {
	c.preset = make(map[string]interface{})
}

// Separator return a current value that is used to split a composite key.
//
// Default '.'
func (c *Config) Separator() string {
	return c.separator
}

// SetSeparator sets a value of a composite key separator.
func (c *Config) SetSeparator(sep string) {
	c.separator = sep
}

// Assemble loads all the data from loaders and merges them with overwriting
// in a order of loaders.
func (c *Config) Assemble() {
	c.conf = c.assemble()
}

func (c *Config) assemble() (result map[string]interface{}) {
	result = c.preset

	for _, loader := range c.loaders {
		confPart, err := loader.GetMap()

		if err != nil {
			// TODO: Add error
			continue
		}

		result = mapop.Merge(result, confPart)
		return
	}

	return
}

// Get return a value getted by a key.
//
// The key can be a composite. Example: 'my.composite.key'.
// The separator can be set using "SetSeparator".
//
// Default: '.'
func (c Config) Get(key string) interface{} {
	value, _ := DeepGet(c.conf, key, c.separator)
	return value
}

// GetAll return all configuration values.
func (c Config) GetAll() map[string]interface{} {
	return c.conf
}

// GetBool return a boolean value obtained by a key.
func (c *Config) GetBool(key string) bool {
	return c.Get(key).(bool)
}

// GetString return a string value obtained by a key.
func (c *Config) GetString(key string) string {
	return c.Get(key).(string)
}

// GetInt return a int value obtained by a key.
func (c *Config) GetInt(key string) int {
	return c.Get(key).(int)
}

// GetUint return a uint value obtained by a key.
func (c *Config) GetUint(key string) uint {
	return c.Get(key).(uint)
}

// GetFloat return a float32 value obtained by a key.
func (c *Config) GetFloat(key string) float32 {
	return c.Get(key).(float32)
}

// GetFloat64 return a float64 value obtained by a key.
func (c *Config) GetFloat64(key string) float64 {
	return c.Get(key).(float64)
}

func (c *Config) Unmarshal(s interface{}) error {
	return mapstructure.Decode(c.conf, s)
}

func (c *Config) UnmarshalKey(key string, s interface{}) error {
	value := c.Get(key)
	return mapstructure.Decode(value, s)
}
