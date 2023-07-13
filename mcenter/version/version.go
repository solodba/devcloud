package version

import "fmt"

// 定义版本信息
var (
	GIT_TAG    string
	GIT_COMMIT string
	GIT_BRANCH string
	BUILD_TIME string
	GO_VERSION string
)

// 显示长版本
func LongVersion() string {
	return fmt.Sprintf("GIT_TAG=%s\nGIT_COMMIT=%s\nGIT_BRANCH=%s\nBUILD_TIME=%s\nGO_VERSION=%s\n",
		GIT_TAG, GIT_COMMIT, GIT_BRANCH, BUILD_TIME, GO_VERSION)
}

// 显示短版本
func ShortVersion() string {
	commit := ""
	if len(GIT_COMMIT) > 8 {
		commit = GIT_COMMIT[:8]
	} else {
		commit = GIT_COMMIT
	}
	return fmt.Sprintf("%s[%s %s]", GIT_TAG, BUILD_TIME, commit)
}
