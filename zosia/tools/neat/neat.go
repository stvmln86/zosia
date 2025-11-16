// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"time"
)

// trim returns a string trimmed to a maximum length.
func trim(text string, size int) string {
	if len(text) <= size {
		return text
	}

	return text[:size]
}

// Body returns a whitespace-stripped body string.
func Body(body string, size int) string {
	body = strings.TrimSpace(body)
	return trim(body, size)
}

// Name returns a lowercase whitespace-stripped name string.
func Name(name string, size int) string {
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	return trim(name, size)
}

// Path returns a cleaned file path string.
func Path(path string) string {
	path = strings.TrimSpace(path)
	return filepath.Clean(path)
}

// Time returns a local Time object from a UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}

// UUID returns a lowercase whitespace-stripped UUID string.
func UUID(uuid string) string {
	uuid = strings.ToLower(uuid)
	uuid = strings.TrimSpace(uuid)
	return trim(uuid, 16)
}
