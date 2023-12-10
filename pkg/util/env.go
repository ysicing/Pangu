// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

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
