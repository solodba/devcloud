package tools

import "github.com/solodba/devcloud/tree/main/mcenter/common/format"

func MustToJson(v any) string {
	return format.PrettifyJson(v)
}
