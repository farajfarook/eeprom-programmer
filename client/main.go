package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := uploadCmd()
	cmd.AddCommand(listPortCmd())
	cmd.AddCommand(hexdumpCmd())
	cmd.AddCommand(versionCmd())
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
