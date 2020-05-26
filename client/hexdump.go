package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func hexdumpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hexdump",
		Short: "Hex Dump the binary file",
		Run: func(cmd *cobra.Command, args []string) {
			file := cmd.Flags().Lookup("file").Value.String()
			if validateFile(file) {
				data, err := ioutil.ReadFile(file)
				if err != nil {
					log.Fatal(err)
				} else {
					hexdump(data, 16)
				}
			}
		},
	}
	cmd.Flags().StringP("file", "f", "", "Binary file")
	cmd.MarkFlagRequired("file")
	viper.BindPFlag("file", cmd.Flags().Lookup("file"))
	return cmd
}

func hexdump(data []byte, bytesPerRow int) {
	datalength := len(data)
	dataPrinted := false
	for pos := 0; pos < datalength; pos += bytesPerRow {
		endPos := pos + bytesPerRow
		if endPos >= datalength {
			endPos = datalength
		}
		rowData := data[pos:endPos]
		prevRowData := []byte{}
		if pos-bytesPerRow >= 0 {
			prevRowData = data[pos-bytesPerRow : pos]
		}
		cmp := bytes.Compare(rowData, prevRowData)
		if pos >= datalength-bytesPerRow || cmp != 0 {
			fmt.Printf("\033[1;34m%08x:\033[0m\t", pos)
			hexdumpln(rowData, bytesPerRow)
			dataPrinted = true
		} else if dataPrinted {
			fmt.Println("\033[1;34m...\033[0m")
			dataPrinted = false
		}
	}
}

func hexdumpln(data []byte, bytesPerRow int) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		fmt.Printf("%02x  ", data[i])
	}
	for i := 0; i < bytesPerRow-dataLen; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("\t\033[1;32m%s\033[0m", viewString(data))
	fmt.Println()
}

func viewString(b []byte) string {
	r := []rune(string(b))
	for i := range r {
		if r[i] < 32 || r[i] > 126 {
			r[i] = '.'
		}
	}
	return string(r)
}
