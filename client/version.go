package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of eeprom-client",
		Long:  `All software has versions. This is EEPROM Client's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("EEPROM Client v0.1.0")
		},
	}
}
