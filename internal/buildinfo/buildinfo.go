// Package buildinfo exposes metadata shared by binaries in this module.
package buildinfo

var version = "dev"

// Version returns the version assigned to this build.
func Version() string {
	return version
}
