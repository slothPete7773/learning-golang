package main

import (
	"context"
	"fmt"
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
	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	// V1: Without Options
	// if err := bucket.Create(ctx, projectId, nil); err != nil {
	// 	log.Fatalf("Failed to create bucket: %v", err)
	// }

	// V2: With Options
	bucketAttrs := storage.BucketAttrs{
		PredefinedACL: "private",
		Location:      "ASIA-SOUTHEAST1",
	}
	if err := bucket.Create(ctx, projectId, &bucketAttrs); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)

}
