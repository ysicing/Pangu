package common

import "fmt"

func GetVersion() string {
	return fmt.Sprintf("%s-%s-%s", Version, BuildDate, GitCommitHash)
}

func GetDefaultPath() string {
	return "/conf/pangu.yaml"
}
