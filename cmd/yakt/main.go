package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "yakt",
	Short:   "yakt: an elegant toolkit for Cloudwego kitex and hertz",
	Long:    "yakt: an elegant toolkit for Cloudwego kitex and hertz",
	Version: version,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
