package handler

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// UploadHandler 上传handler
func UploadHandler(ch chan<- string, tmpl *template.Template) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				err := r.ParseMultipartForm(32 << 20)
				if err != nil {
					log.Fatal(err)
				}
				f, header, err := r.FormFile("file")
				if err != nil {
					log.Fatal(err)
				}
				filename := header.Filename
				file, err := os.Create(filename)
				_, err = io.Copy(file, f)
				if err != nil {
					log.Fatal(err)
				}
				ch <- filename
			} else {
				tmpl.Execute(w, nil)
			}
		})
}

// DownloadHandler 下载handler
func DownloadHandler(filename string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open(filename)
			if err != nil {
				http.Error(w, "File not found", http.StatusNotFound)
			}
			defer f.Close()

			// 	http://mrwaggel.be/post/golang-transmit-files-over-a-nethttp-server-to-clients/
			//File is found, create and send the correct headers

			//Get the Content-Type of the file
			//Create a buffer to store the header of the file in
			FileHeader := make([]byte, 512)
			//Copy the headers into the FileHeader buffer
			f.Read(FileHeader)
			//Get content type of file
			FileContentType := http.DetectContentType(FileHeader)

			//Get the file size
			FileStat, _ := f.Stat()                            //Get info from file
			FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

			//Send the headers
			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Type", FileContentType)
			w.Header().Set("Content-Length", FileSize)

			//Send the file
			//We read 512 bytes from the file already, so we reset the offset back to 0
			f.Seek(0, 0)
			_, err = io.Copy(w, f) //'Copy' the file to the client
			if err != nil {
				log.Fatal(err)
			}
		})
}
