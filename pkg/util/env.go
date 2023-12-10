package util

import (
	"github.com/spf13/viper"
)

func GetKeyFromYaml(key string, defaultvalue ...string) string {
	dk := ""
	if len(defaultvalue) > 0 {
		dk = defaultvalue[0]
	}
	getKey := viper.GetString(key)
	if len(getKey) == 0 {
		return dk
	}
	return getKey
}

func GetKeyIntFromYaml(key string, defaultvalue ...int) int {
	dk := 0
	if len(defaultvalue) > 0 {
		dk = defaultvalue[0]
	}
	getKey := viper.GetInt(key)
	if getKey == 0 {
		return dk
	}
	return getKey
}

func GetStatusFromYaml(key string) bool {
	return viper.GetBool(key)
}
