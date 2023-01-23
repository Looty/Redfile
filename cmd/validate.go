/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	config "redfile/internal/config/utils"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate redfile config file against the scheme",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

redfile validate <config_path>`,
	Run: func(cmd *cobra.Command, args []string) {
		validate()
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&File, "file", "f", "redfile.yaml", "Redfile to read from")
}

func validate() {
	config.ValidateRedfile(File)
}
