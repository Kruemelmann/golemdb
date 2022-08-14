package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Create or update a Key",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
