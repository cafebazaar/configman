package configman

import (
	"time"

	"github.com/fsnotify/fsnotify"
)

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

	SetEnvPrefix(string)
	// RegisterAlias provides another accessor for the same key.
	// This enables one to change a name without breaking the application
	RegisterAlias(alias string, key string)

	// ReadInConfig will discover and load the configuration file from disk
	// and key/value stores, searching in one of the defined paths.
	ReadInConfig() error

	SetConfigFile(string)

	WatchConfig() error
	OnConfigChange(func(fsnotify.Event))
}
