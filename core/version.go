package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// VersionInfo contains version details.
type VersionInfo struct {
	VersionMajor int    `json:"VersionMajor"`
	VersionMinor int    `json:"VersionMinor"`
	VersionPatch int    `json:"VersionPatch"`
	VersionMeta  string `json:"VersionMeta"`
}

// readVersionInfo reads version information from version.json.
func readVersionInfo() VersionInfo {
	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Find the 'razor-go' root directory.
	rootDir, err := findRazorGoRoot(cwd)
	if err != nil {
		log.Fatalf("Error finding 'razor-go' root directory: %v", err)
	}

	// Construct the path to the version.json file based on the executable location.
	versionFilePath := filepath.Join(rootDir, "version.json")
	versionJSONFile, err := os.Open(versionFilePath)
	if err != nil {
		log.Fatalf("Error in opening version.json file: %v", err)
	}
	defer versionJSONFile.Close() // Ensure file is closed after reading

	data, err := io.ReadAll(versionJSONFile)
	if err != nil {
		log.Fatalf("Error reading version.json: %v", err)
	}

	var v VersionInfo
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatalf("Error unmarshalling version.json: %v", err)
	}
	return v
}

// VersionWithMeta returns the textual version string including the metadata.
func VersionWithMeta() string {
	v := readVersionInfo()
	version := fmt.Sprintf("%d.%d.%d", v.VersionMajor, v.VersionMinor, v.VersionPatch)
	if v.VersionMeta != "" {
		version += "-" + v.VersionMeta
	}
	return version
}

// findRazorGoRoot attempts to find the 'razor-go' directory by walking up the hierarchy.
func findRazorGoRoot(startDir string) (string, error) {
	for {
		if filepath.Base(startDir) == "razor-go" {
			return startDir, nil
		}

		parentDir := filepath.Dir(startDir)
		if parentDir == startDir {
			// No more directories above, we've reached the root of the file system.
			return "", fmt.Errorf("'razor-go' directory not found in path hierarchy")
		}

		startDir = parentDir
	}
}
