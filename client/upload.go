package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func uploadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload binary to host",
		Run: func(cmd *cobra.Command, args []string) {
			file := cmd.Flags().Lookup("file").Value.String()
			port := cmd.Flags().Lookup("port").Value.String()
			if validateFile(file) && validatePort(port) {
				upload(file, port)
			}
		},
	}
	cmd.Flags().StringP("file", "f", "", "Binary file to upload")
	cmd.MarkFlagRequired("file")
	cmd.Flags().StringP("port", "p", "", "Serial port attached")
	cmd.MarkFlagRequired("port")
	viper.BindPFlag("file", cmd.Flags().Lookup("file"))
	viper.BindPFlag("port", cmd.Flags().Lookup("port"))

	return cmd
}

func upload(file string, port string) {
	fmt.Println("Uploading binary file " + file + " to port " + port)
	hexdump(([]byte)("Hello how are you.. sdfsdfdsf sd fds f asd"), 16)
}
