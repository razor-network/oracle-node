package core

import "testing"

// TestVersionWithMeta tests the VersionWithMeta function using the actual version.json file.
func TestVersionWithMeta(t *testing.T) {
	expectedVersion := "1.1.0" // This should match the contents of version.json

	version := VersionWithMeta()
	if version != expectedVersion {
		t.Errorf("VersionWithMeta() = %v; want %v", version, expectedVersion)
	}
}
