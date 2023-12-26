package main

import (
	"fmt"
	"net/http"

	"github.com/01zulfi/file-uploader/controllers"
)

const port = "8080"

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", controllers.HandleIndex)
	server.HandleFunc("/upload", controllers.HandleUpload)
	server.HandleFunc("/download/", controllers.HandleDownload)

	fmt.Println("server started at http://localhost:" + port)
	err := http.ListenAndServe(":"+port, server)

	if err != nil {
		fmt.Println("error while starting server")
	}
}
