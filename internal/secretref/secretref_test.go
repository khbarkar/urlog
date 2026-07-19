package secretref

import (
	"errors"
	"testing"
)

func TestParseValidReference(t *testing.T) {
	ref, err := Parse(" secret://aws-secrets-manager/eu-north-1/urlog/forgeboard/github-token ")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if ref.Backend != "aws-secrets-manager" {
		t.Fatalf("Backend = %q, want aws-secrets-manager", ref.Backend)
	}
	if ref.Path != "eu-north-1/urlog/forgeboard/github-token" {
		t.Fatalf("Path = %q, want eu-north-1/urlog/forgeboard/github-token", ref.Path)
	}
	if got := ref.String(); got != "secret://aws-secrets-manager/eu-north-1/urlog/forgeboard/github-token" {
		t.Fatalf("String() = %q", got)
	}
}

func TestParseRejectsInvalidReferences(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want error
	}{
		{name: "empty", raw: " ", want: ErrEmptyReference},
		{name: "missing scheme", raw: "aws-secrets-manager/urlog/token", want: ErrInvalidScheme},
		{name: "missing backend", raw: "secret:///urlog/token", want: ErrMissingBackend},
		{name: "missing path", raw: "secret://aws-secrets-manager", want: ErrMissingPath},
		{name: "slash only path", raw: "secret://aws-secrets-manager///", want: ErrMissingPath},
		{name: "password assignment", raw: "password=correct-horse-battery-staple", want: ErrSuspiciousCleartext},
		{name: "bearer token", raw: "Bearer abc.def.ghi", want: ErrSuspiciousCleartext},
		{name: "private key", raw: "-----BEGIN PRIVATE KEY-----", want: ErrSuspiciousCleartext},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.raw)
			if !errors.Is(err, tt.want) {
				t.Fatalf("Parse(%q) error = %v, want %v", tt.raw, err, tt.want)
			}
		})
	}
}

func TestLooksValid(t *testing.T) {
	if !LooksValid("secret://local-sops-age/forgeboard/github-token") {
		t.Fatal("LooksValid returned false for valid reference")
	}
	if LooksValid("token=cleartext") {
		t.Fatal("LooksValid returned true for cleartext")
	}
}

func TestMustParsePanicsOnInvalidReference(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustParse did not panic")
		}
	}()
	MustParse("not-a-secret-ref")
}
