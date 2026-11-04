package buildinfo_test

import (
	"testing"

	"github.com/Amal-Abushiesha/amal-abushiesha-dark-factory/internal/buildinfo"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{name: "default version", want: "dev"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := buildinfo.Version(); got != test.want {
				t.Errorf("Version() = %q, want %q", got, test.want)
			}
		})
	}
}
