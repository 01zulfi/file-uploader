package storage

import (
	"fmt"
	"os"

	supabaseStorage "github.com/supabase-community/storage-go"
)

var storageClient *supabaseStorage.Client

const BucketName = "file_uploader"

func Init() error {
	storageUrl, ok := os.LookupEnv("STORAGE_URL")
	if !ok {
		return fmt.Errorf("STORAGE_URL environment variable is not set")
	}

	storageToken, ok := os.LookupEnv("STORAGE_TOKEN")
	if !ok {
		return fmt.Errorf("STORAGE_TOKEN environment variable is not set")
	}

	storageClient = supabaseStorage.NewClient(storageUrl, storageToken, nil)
	return nil
}

func Get() *supabaseStorage.Client {
	return storageClient
}
