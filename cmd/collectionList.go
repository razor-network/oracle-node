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
	cmdUtilsMockery.ExecuteCollectionList()

}

func (*UtilsStructMockery) ExecuteCollectionList() {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	err = cmdUtilsMockery.GetCollectionList(client)
	utils.CheckError("Error in getting collection list: ", err)
}

func (*UtilsStructMockery) GetCollectionList(client *ethclient.Client) error {
	collections, err := razorUtilsMockery.GetCollections(client)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Active", "Collection Id", "Asset Index", "Power", "Aggregation Method", "Job IDs", "Name", "Tolerance"})
	for i := 0; i < len(collections); i++ {
		jobIDs, _ := json.Marshal(collections[i].JobIDs)

		table.Append([]string{
			strconv.FormatBool(collections[i].Active),
			strconv.Itoa(int(collections[i].Id)),
			strconv.Itoa(int(collections[i].AssetIndex)),
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
	razorUtilsMockery = UtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

	rootCmd.AddCommand(collectionListCmd)

}
