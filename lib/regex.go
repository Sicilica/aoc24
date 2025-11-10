package lib

import "regexp"

func Match(r *regexp.Regexp, s string) []string {
	m := r.FindStringSubmatch(s)
	Assert(m != nil)
	return m
}
