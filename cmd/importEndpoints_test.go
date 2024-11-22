package cmd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"razor/core"
	"testing"
)

func TestExecuteImportEndpoints_WithMockedPath(t *testing.T) {
	// Create a temporary directory to mock the default path
	tempDir, err := os.MkdirTemp("", "test_default_path")
	assert.NoError(t, err, "Temporary directory should be created")
	defer os.RemoveAll(tempDir) // Clean up after the test

	SetUpMockInterfaces()

	pathMock.On("GetDefaultPath").Return(tempDir, nil)

	// Call ExecuteImportEndpoints
	utils := &UtilsStruct{}
	utils.ExecuteImportEndpoints()

	// Verify the file was created in the mocked directory
	destFilePath := filepath.Join(tempDir, "endpoints.json")
	_, err = os.Stat(destFilePath)
	assert.NoError(t, err, "endpoints.json should be created in the mocked directory")

	// Verify the content of the file matches core.DefaultEndpoints
	data, err := os.ReadFile(destFilePath)
	assert.NoError(t, err, "File should be readable")

	var importedEndpoints []string
	err = json.Unmarshal(data, &importedEndpoints)
	assert.NoError(t, err, "JSON should be valid")
	assert.Equal(t, core.DefaultEndpoints, importedEndpoints, "Imported endpoints should match the defaults")

	// Delete the file after verification
	err = os.Remove(destFilePath)
	assert.NoError(t, err, "endpoints.json should be deleted successfully")

	// Verify logs
	log.Infof("Default endpoints successfully imported to %s", destFilePath)
	pathMock.AssertExpectations(t)

	// Confirm the file no longer exists
	_, err = os.Stat(destFilePath)
	assert.Error(t, err, "endpoints.json should not exist after deletion")
	assert.True(t, os.IsNotExist(err), "Error should indicate file does not exist")
}
