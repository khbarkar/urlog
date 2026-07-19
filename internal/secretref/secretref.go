// Package secretref validates Urlog secret references.
//
// Urlog config may contain references to secrets, but never secret values.
// This package is intentionally small so bootstrap, Integration, and policy
// code can share one strict interpretation of what a secret reference is.
package secretref

import (
	"errors"
	"fmt"
	"strings"
)

const scheme = "secret://"

var (
	ErrEmptyReference      = errors.New("secret reference is empty")
	ErrInvalidScheme       = errors.New("secret reference must start with secret://")
	ErrMissingBackend      = errors.New("secret reference is missing backend")
	ErrMissingPath         = errors.New("secret reference is missing path")
	ErrSuspiciousCleartext = errors.New("value looks like cleartext, not a secret reference")
)

// Reference is a parsed secret:// reference.
type Reference struct {
	Backend string
	Path    string
}

// Parse validates and parses a secret reference.
func Parse(raw string) (Reference, error) {
	value := strings.TrimSpace(raw)
	if value == "" {
		return Reference{}, ErrEmptyReference
	}
	if looksLikeCleartext(value) {
		return Reference{}, ErrSuspiciousCleartext
	}
	if !strings.HasPrefix(value, scheme) {
		return Reference{}, ErrInvalidScheme
	}

	rest := strings.TrimPrefix(value, scheme)
	backend, path, ok := strings.Cut(rest, "/")
	if !ok {
		return Reference{}, ErrMissingPath
	}
	backend = strings.TrimSpace(backend)
	path = strings.Trim(path, "/")

	if backend == "" {
		return Reference{}, ErrMissingBackend
	}
	if path == "" {
		return Reference{}, ErrMissingPath
	}
	if strings.ContainsAny(backend, " \t\r\n") {
		return Reference{}, fmt.Errorf("secret backend %q contains whitespace", backend)
	}

	return Reference{Backend: backend, Path: path}, nil
}

// MustParse is for static test fixtures and package-level examples.
func MustParse(raw string) Reference {
	ref, err := Parse(raw)
	if err != nil {
		panic(err)
	}
	return ref
}

// String renders the reference back to its canonical form.
func (r Reference) String() string {
	return scheme + r.Backend + "/" + strings.Trim(r.Path, "/")
}

// LooksValid reports whether raw is a valid secret reference.
func LooksValid(raw string) bool {
	_, err := Parse(raw)
	return err == nil
}

func looksLikeCleartext(value string) bool {
	lower := strings.ToLower(value)
	cleartextHints := []string{
		"password=",
		"passwd=",
		"token=",
		"api_key=",
		"apikey=",
		"secret=",
		"bearer ",
		"-----begin ",
	}
	for _, hint := range cleartextHints {
		if strings.Contains(lower, hint) {
			return true
		}
	}
	return false
}
