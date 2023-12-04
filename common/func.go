package common

import "fmt"

func GetVersion() string {
	return fmt.Sprintf("%s-%s-%s", Version, BuildDate, GitCommitHash)
}
