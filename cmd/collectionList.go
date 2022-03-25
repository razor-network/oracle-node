package cmd

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"razor/utils"
	"strconv"
	"strings"
)

// collectionListCmd represents the collectionList command
var collectionListCmd = &cobra.Command{
	Use:   "collectionList",
	Short: "list of all collections",
	Long: `Provides the list of all collections with their name, power, ID etc.
Example:
	./razor collectionList `,
	Run: initialiseCollectionList,
}

func initialiseCollectionList(*cobra.Command, []string) {
	cmdUtils.ExecuteCollectionList()
}

func (*UtilsStruct) ExecuteCollectionList() {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	err = cmdUtils.GetCollectionList(client)
	utils.CheckError("Error in getting collection list: ", err)
}

func (*UtilsStruct) GetCollectionList(client *ethclient.Client) error {
	collections, err := razorUtils.GetCollections(client)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Active", "Collection Id", "Power", "Aggregation Method", "Job IDs", "Name", "Tolerance"})
	for i := 0; i < len(collections); i++ {
		jobIDs, _ := json.Marshal(collections[i].JobIDs)

		table.Append([]string{
			strconv.FormatBool(collections[i].Active),
			strconv.Itoa(int(collections[i].Id)),
			strconv.Itoa(int(collections[i].Power)),
			strconv.Itoa(int(collections[i].AggregationMethod)),
			strings.Trim(string(jobIDs), "[]"),
			collections[i].Name,
			strconv.Itoa(int(collections[i].Tolerance)),
		})

	}

	table.Render()
	return nil

}

func init() {
	rootCmd.AddCommand(collectionListCmd)

}
