package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.bug.st/serial"
)

func listPortCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "listport",
		Short: `List available Serial ports`,
		Run: func(cmd *cobra.Command, args []string) {
			ports := listPorts()
			for _, portname := range ports {
				fmt.Println(portname)
			}
		},
	}
}

func listPorts() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	return ports
}

func validatePort(port string) bool {
	ports := listPorts()
	for _, portname := range ports {
		if portname == port {
			return true
		}
	}
	log.Fatal("Invalid serial port " + port)
	return false
}
