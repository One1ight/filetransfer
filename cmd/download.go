/*
Copyright © 2020 One_1ight <One_1ight@hotmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"transfer/handler"
	"transfer/qrcode"
	"transfer/server"
	"transfer/utils"

	"github.com/spf13/cobra"
	"rsc.io/qr"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "指定文件名，扫码下载。完成后按ctrl+c结束",
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == "" {
			return errors.New("filename is empty")
		}
		port := ":8888"
		filename := args[0]
		qrcode.Generate("http://"+utils.GetLocalIP()+port, qr.M, os.Stdout)
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)
		sc, err := server.ListenAndServeWithClose(":8888", handler.DownloadHandler(filename))
		if err != nil {
			return err
		}
		select {
		case <-ch:
			// Close HTTP Server
			err := sc.Close()
			if err != nil {
				log.Fatalln("Server Close Error - ", err)
			}
			log.Println("Server Closed")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
