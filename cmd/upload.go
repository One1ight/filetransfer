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
	"html/template"
	"log"
	"os"
	"os/exec"
	"transfer/handler"
	"transfer/qrcode"
	"transfer/server"
	"transfer/utils"

	"github.com/spf13/cobra"
	"rsc.io/qr"
)

var tmpl *template.Template

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "扫码后，打开页面并指定文件上传。",
	Run: func(cmd *cobra.Command, args []string) {
		port := ":8888"
		// 阻塞&传递filename
		ch := make(chan string)
		srvCLoser, err := server.ListenAndServeWithClose(port, handler.UploadHandler(ch, tmpl))
		if err != nil {
			log.Fatalln("ListenAndServeWithClose Error - ", err)
		}
		// 将地址
		qrcode.Generate("http://"+utils.GetLocalIP()+port, qr.M, os.Stdout)
		// post finished?
		select {
		case filename := <-ch:
			// Close HTTP Server
			err := srvCLoser.Close()
			if err != nil {
				log.Fatalln("Server Close Error - ", err)
			}
			log.Println("Server Closed")
			cmd := exec.Command("open", filename)
			cmd.Run()
		}

	},
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmp/form.html"))
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
