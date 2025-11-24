package api

import "net/url"

// parseQuery parses query parameters and returns an error if any malformed value pairs are encountered.
func parseQuery(rawQuery string) (url.Values, error) {
	return url.ParseQuery(rawQuery)
}
