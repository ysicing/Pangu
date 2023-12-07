// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package main

import (
	"runtime"

	"gitea.ysicing.net/cloud/pangu/cmd"
	errors "github.com/ergoapi/util/exerror"
)

// @title Pangu API
// @version 0.0.1
// @description pangu.

// @contact.name ysicing
// @contact.url http://github.com/ysicing/pangu
// @contact.email i@ysicing.me

// @license.name AGPLv3
// @license.url https://opensource.org/licenses/MIT
func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	errors.CheckAndExit(cmd.Execute())
}
