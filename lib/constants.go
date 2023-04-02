package lib

import (
	"regexp"
)

// REGEX for checking website URL.
var WEBSITE_REGEX = regexp.MustCompile(`(?i)^(https:\/\/|http:\/\/|www.){1}[^-][a-zA-Z0-9-]{1,50}\.[^-][a-zA-Z0-9()]{1,50}\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)$`)
