package cmd

import (
	"fmt"
	"foodies/gosrc/server"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "foodies",
	Short: "Search food trucks in SF",
	Long:  `Search food trucks in SF`,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

// addCmd represents the add command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Long:  `Starts foodies server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
