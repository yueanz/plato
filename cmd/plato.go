package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plato",
	Short: "plato short",
	Long:  "plato long",
	Run:   plato,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func plato(cmd *cobra.Command, args []string) {
	fmt.Println("hello world v2")
}
