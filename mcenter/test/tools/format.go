package tools

import "github.com/solodba/mcube/format"

func MustToJson(v any) string {
	return format.PrettifyJson(v)
}
