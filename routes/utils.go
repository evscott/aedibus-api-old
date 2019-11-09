package Routes

func Path(resources ...Resource) string {
	route := ""
	for _, r := range resources {
		route += string(r)
	}
	return route
}

func String(s string) *string {
	return &s
}
