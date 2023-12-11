// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package common

import (
	"fmt"

	"github.com/ergoapi/util/file"
	"github.com/ergoapi/util/zos"
)

func GetVersion() string {
	return fmt.Sprintf("%s-%s-%s", Version, BuildDate, GitCommitHash)
}

func GetDefaultPath() string {
	return "/conf/pangu.yaml"
}

func GetDefaultLogFile() string {
	logfile := "/tmp/pangu.log"
	if zos.IsLinux() {
		logfile = "/var/log/pangu/pangu.log"
		file.MkFileFullPathDir(logfile)
	}
	return logfile
}
