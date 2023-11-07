package env

import (
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func getString(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if value := viper.GetString(key); value != "" {
		return value
	}
	return ""
}

func getInt(key string) *int {
	if value := getString(key); value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return &v
		}
	}
	return nil
}

func getBool(key string) *bool {
	if value := getString(key); value != "" {
		if v, err := strconv.ParseBool(value); err == nil {
			return &v
		}
	}
	return nil
}

func getFloat64(key string) *float64 {
	if value := getString(key); value != "" {
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			return &v
		}
	}
	return nil
}

// GetString return the value as string from the environment variable or the value of the key in the config file.
func GetString(key string, defaultValue ...string) string {
	if value := getString(key); value != "" {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetInt returns the value as int from the environment variable or the value of the key in the config file.
func GetInt(key string, defaultValue ...int) int {
	if value := getInt(key); value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetBool return the value as bool from the environment variable or the value of the key in the config file.
func GetBool(key string, defaultValue ...bool) bool {
	if value := getBool(key); value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

// GetFloat64 returns the value as float64 from the environment variable or the value of the key in the config file.
func GetFloat64(key string, defaultValue ...float64) float64 {
	if value := getFloat64(key); value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetDurationMs returns the value as time.Duration from the environment variable or the value of the key in the config file.
func GetDurationMs(key string, defaultValue ...int) time.Duration {
	value := GetInt(key)
	if value == 0 && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return time.Millisecond * time.Duration(value)
}

// GetDurationSeconds returns the value as time.Duration from the environment variable or the value of the key in the config file.
func GetDurationSeconds(key string, defaultValue ...float64) time.Duration {
	value := GetFloat64(key)
	if value == 0 && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return time.Duration(value * float64(time.Second))
}

// GetDurationMinutes returns the value as time.Duration from the environment variable or the value of the key in the config file.
func GetDurationMinutes(key string, defaultValue ...float64) time.Duration {
	value := GetFloat64(key)
	if value == 0 && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return time.Duration(value * float64(time.Minute))
}

// MustGetString return the value as string from the environment variable or the value of the key in the config file.
//
// If the value is empty, it will panic.
func MustGetString(key string) string {
	if value := getString(key); value != "" {
		return value
	}
	panic("env: " + key + " is not set")
}

// MustGetInt returns the value as int from the environment variable or the value of the key in the config file.
//
// If the value is empty, it will panic.
func MustGetInt(key string) int {
	if value := getInt(key); value != nil {
		return *value
	}
	panic("env: " + key + " is not set")
}

// MustGetBool return the value as bool from the environment variable or the value of the key in the config file.
//
// If the value is empty, it will panic.
func MustGetBool(key string) bool {
	if value := getBool(key); value != nil {
		return *value
	}
	panic("env: " + key + " is not set")
}

// MustGetFloat64 returns the value as float64 from the environment variable or the value of the key in the config file.
//
// If the value is empty, it will panic.
func MustGetFloat64(key string) float64 {
	if value := getFloat64(key); value != nil {
		return *value
	}
	panic("env: " + key + " is not set")
}

// MustGetDurationMs returns the value as time.Duration from the environment variable or the value of the key in the config file.
//
// The value is in milliseconds, if the value is empty, it will panic.
func MustGetDurationMs(key string) time.Duration {
	return time.Millisecond * time.Duration(MustGetInt(key))
}

// MustGetDurationSeconds returns the value as time.Duration from the environment variable or the value of the key in the config file.
//
// The value is in seconds, if the value is empty, it will panic.
func MustGetDurationSeconds(key string) time.Duration {
	return time.Duration(MustGetFloat64(key) * float64(time.Second))
}

// MustGetDurationMinutes returns the value as time.Duration from the environment variable or the value of the key in the config file.
//
// The value is in minutes, if the value is empty, it will panic.
func MustGetDurationMinutes(key string) time.Duration {
	return time.Duration(MustGetFloat64(key) * float64(time.Minute))
}
