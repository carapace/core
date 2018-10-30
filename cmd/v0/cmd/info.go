package cmd

import (
	"context"
	"fmt"
	"github.com/carapace/core/cmd/v0/carapace"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Returns information on the current binary configuration",
	Long:  `info returns information on the current binary, it's configured handlers, api versions and the node's configuration status.`,
	Run: func(cmd *cobra.Command, args []string) {
		app, _, err := carapace.New(logLevel)

		if err != nil {
			fmt.Println("internal error occurred while initializing server: ", err.Error())
		}

		info, err := app.InfoService(context.Background(), nil)
		if err != nil {
			fmt.Println("internal error occurred while initializing server: ", err.Error())
		}

		fmt.Println(fmt.Sprintf("%+v\n", info))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
