/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a Cover Letter pdf file",
	Long: `Generates a Cover Letter pdf file based on a company name, position, and a default template.
	To use the command pass in the company name as the first argument and the position as the second argument. 
	Example: cover-letter-generator generate Google "Senior Engineer" (Note if the position is two words pass it in with quotes)
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating Cover Letter for the position %s at company %s\n", args[0], args[1])
		err := GenerateCoverLetter(args[0], args[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
