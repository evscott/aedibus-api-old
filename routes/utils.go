package Routes

import (
	"strings"
)

func Path(resources ...Resource) string {
	var route strings.Builder
	for _, r := range resources {
		route.WriteString(string(r))
	}

	return route.String()
}

func String(s string) *string {
	return &s
}
