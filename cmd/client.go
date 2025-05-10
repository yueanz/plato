package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yueanz/plato/client"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use: "client",
	Run: ClientHandler,
}

func ClientHandler(cmd *cobra.Command, args []string) {
	fmt.Println("client")
	client.RunMain()
}
