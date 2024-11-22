package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"razor/core"
	"razor/utils"
)

// importEndpointsCmd represents the importEndpoints command
var importEndpointsCmd = &cobra.Command{
	Use:   "importEndpoints",
	Short: "imports the list of endpoints locally",
	Long: `Imports the list of endpoints to $HOME/.rzr directory allowing the user to access/edit it easily.
Example:
	./razor importEndpoints `,
	Run: initialiseImportEndpoints,
}

func initialiseImportEndpoints(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteImportEndpoints()
}

// ExecuteImportEndpoints imports the list of endpoints from core/endpoints.go to $HOME/.razor directory locally.
func (*UtilsStruct) ExecuteImportEndpoints() {
	defaultPath, err := pathUtils.GetDefaultPath()
	utils.CheckError("Error in getting default path: ", err)

	// Define the target path for endpoints.json
	destFilePath := filepath.Join(defaultPath, "endpoints.json")

	// Serialize the default endpoints to JSON
	endpointsData, err := json.MarshalIndent(core.DefaultEndpoints, "", "  ")
	utils.CheckError("Error in serializing the endpoints: %w", err)

	// Write the JSON to the destination file
	err = os.WriteFile(destFilePath, endpointsData, 0644)
	utils.CheckError("Error in writing endpoints.json: %w", err)

	log.Infof("Default endpoints successfully imported to %s", destFilePath)
}

func init() {
	rootCmd.AddCommand(importEndpointsCmd)
}
