package core

import "fmt"

const (
	VersionMajor = 1                                    // Major version component of the current release
	VersionMinor = 0                                    // Minor version component of the current release
	VersionPatch = 3                                    // Patch version component of the current release
	VersionMeta  = "incentivised-testnet-phase2-patch3" // Version metadata to append to the version string
)

// Version holds the textual version string.
var Version = func() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}()

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := Version
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()
