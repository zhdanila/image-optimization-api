package schema

import "strings"

func ParentIsMissing(err error) bool {
	return strings.Contains(err.Error(), `is missing. Row cannot`)
}
