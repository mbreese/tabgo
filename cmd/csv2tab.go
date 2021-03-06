package cmd

import (
	"fmt"
	"os"

	"github.com/mbreese/tabl/textfile"
	"github.com/spf13/cobra"
)

func init() {
	csv2TabCmd.Flags().BoolVarP(&ShowComments, "show-comments", "H", false, "Show comments")
	rootCmd.AddCommand(csv2TabCmd)
}

var csv2TabCmd = &cobra.Command{
	Use:   "csv2tab",
	Short: "Convert a CSV file to tab-delimited format",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && args[0] != "-" {
			_, err := os.Stat(args[0])
			if os.IsNotExist(err) {
				return fmt.Errorf("Missing file: %s", args[0])
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"-"}
		}
		txt := textfile.NewCSVFile(args[0]).
			WithNoHeader(true)

		textfile.NewCSVExporter(txt).
			WithShowComments(ShowComments).
			WriteFile(os.Stdout)
	},
}
