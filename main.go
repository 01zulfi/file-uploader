package main

import (
	"fmt"
	"net/http"

	"github.com/01zulfi/file-uploader/controllers"
	"github.com/01zulfi/file-uploader/db"
)

const (
	port = "8080"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println("error while initializing database")
		return
	}

	server := http.NewServeMux()
	server.HandleFunc("/", controllers.Get(controllers.HandleIndex))
	server.HandleFunc("/login", controllers.GetOrPost(controllers.HandleLogin))
	server.HandleFunc("/logout", controllers.Get(controllers.HandleLogout))
	server.HandleFunc("/upload", controllers.Post(controllers.HandleUpload))
	server.HandleFunc("/download/", controllers.Get(controllers.HandleDownload))

	fmt.Println("server started at http://localhost:" + port)
	err = http.ListenAndServe(":"+port, server)

	if err != nil {
		fmt.Println("error while starting server")
	}
}
