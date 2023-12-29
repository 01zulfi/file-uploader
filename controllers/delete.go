package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

type DeleteRequest struct {
	Filepath string `json:"filepath"`
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	var deleteReq DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&deleteReq)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = data.DeleteFile(deleteReq.Filepath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
