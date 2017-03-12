package configman

import "time"

type ConfigManager interface {
	Get(string) interface{}
	GetString(string) string
	GetBool(string) bool
	GetInt(string) int
	GetInt64(string) int64
	GetFloat64(string) float64
	GetTime(string) time.Time
	GetDuration(string) time.Duration
	GetStringSlice(string) []string
	GetStringMap(string) map[string]interface{}
	GetStringMapString(string) map[string]string
	GetStringMapStringSlice(string) map[string][]string
	GetSizeInBytes(string) uint

	Set(key string, value interface{})

	IsSet(key string) bool
	InConfig(key string) bool
	AllKeys() []string
	SetDefault(key string, value interface{})

	// Sub returns new Config instance representing a sub tree of this instance.
	Sub(key string) ConfigManager

	// Unmarshal unmarshals the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(rawVal interface{}) error
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	UnmarshalKey(key string, rawVal interface{}) error

	// RegisterAlias provides another accessor for the same key.
	// This enables one to change a name without breaking the application
	RegisterAlias(alias string, key string)

	WatchConfig() error
	AddToChangeListeners(ConfigChangeListener)
	RemoveFromChangeListeners(ConfigChangeListener)
}

type ConfigChangeListener interface {
	OnConfigChanged()
}
