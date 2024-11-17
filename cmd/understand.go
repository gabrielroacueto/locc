/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gabrielroacueto/locc/api"
	"github.com/spf13/cobra"
)

// understandCmd represents the understand command
var understandCmd = &cobra.Command{
	Use:   "understand",
	Short: "Given a directory for a repository, understand what it does.",
	Long:  `Use to understand repository. It will prompt an LLM with directory structure and try to figure out what's up.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]

		callback := func(chunk string) {
			fmt.Print(chunk)
		}

		err := api.StreamDirectoryAnalysis(directory, callback)

		fmt.Printf("Error when trying to stream directory analysys: %s\n", err)
	},
}

func init() {
	rootCmd.AddCommand(understandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// understandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// understandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}