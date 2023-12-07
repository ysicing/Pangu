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
// @contact.url http://github.com/ysicing
// @contact.email i@ysicing.me

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	errors.CheckAndExit(cmd.Execute())
}
