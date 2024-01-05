/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/JakubOboza/randal/config"
	"github.com/JakubOboza/randal/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")

		if err != nil {
			fmt.Println(err)
			return
		}

		filePath, err := cmd.Flags().GetString("file")

		if err != nil {
			fmt.Println(err)
			return
		}

		fileContent, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}

		conf, err := config.Load(fileContent)

		if err != nil {
			fmt.Println(err)
			return
		}

		server := server.New(port, conf)
		err = server.Setup()

		if err != nil {
			fmt.Println("Error while initializing randal: ", err)
			return
		}

		fmt.Println("Starting randal...")
		server.Start()

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringP("file", "f", "config.yml", "randal config file location, default name is conifg.yml")
	serveCmd.Flags().IntP("port", "p", 5566, "port on which randal will accept connections. Default port is 5566")
}
