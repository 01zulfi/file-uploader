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
		fmt.Println(err)
		fmt.Println("error while initializing database")
		return
	}

	server := http.NewServeMux()
	server.HandleFunc("/", controllers.Protect(controllers.Get(controllers.HandleIndex)))
	server.HandleFunc("/login", controllers.GetOrPost(controllers.HandleLogin))
	server.HandleFunc("/logout", controllers.Get(controllers.HandleLogout))
	server.HandleFunc("/upload", controllers.Protect(controllers.Post(controllers.HandleUpload)))
	server.HandleFunc("/download/", controllers.Protect(controllers.Get(controllers.HandleDownload)))

	fmt.Println("server started at port:" + port)
	err = http.ListenAndServe("0.0.0.0:8080", server)

	if err != nil {
		fmt.Println("error while starting server")
	}
}
