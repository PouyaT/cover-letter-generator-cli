/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.3"

// TODO: Figure out why the command is working for cover-letter-generator and not cl-gen
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: version,
	Use:     "cover-letter-generator [command] [flags]",
	Short:   "A CLI to generate Cover Letters",
	Long: `Cover Letter Generator is a CLI to create Cover Letters from a default template and will output the cover letter as a pdf.
	To use the command call cover-letter-generator generate Google "Senior Engineer" (Note if the position is two words pass it in with quotes)`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cover-letter-generator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
