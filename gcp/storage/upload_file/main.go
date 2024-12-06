package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../../../credential/gcp-service-account/slothpete7773-warehouse.json")
	ctx := context.Background()

	projectId := "slothpete7773-warehouse"
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("Failed to create client: %v", err)
	}
	defer client.Close()

	bucketName := "new_go_test_bucket"
	storageClient := &StorageClient{
		client:     client,
		projectId:  projectId,
		bucketName: bucketName,
		uploadPath: "/test",
	}

	storageClient.uploadFile(os.Stdout, "./greet.txt", "greet.txt")

	// uploadFile(os.Stdout, "./greet.txt")

}

type StorageClient struct {
	client     *storage.Client
	projectId  string
	bucketName string
	uploadPath string
}

func (c *StorageClient) uploadFile(w io.Writer, srcPath string, objectName string) error {
	ctx := context.Background()

	// Open local file
	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	object := c.client.Bucket(c.bucketName).Object(c.uploadPath + objectName)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.

	// V1: For an object that does not yet exist, set the DoesNotExist precondition.
	object = object.If(storage.Conditions{DoesNotExist: true})

	// V2: If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %w", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := object.NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fmt.Fprintf(w, "Blob %v uploaded.\n", objectName)
	return nil
}
