package main

import (
	"backend/db"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"firebase.google.com/go/v4/storage"
)

func main() {
	ctx := context.Background()
    filePath := "images/test.png"

	storageClient, bucketName, err := db.NewFirebase()
	if err != nil {
		log.Fatalln(err)
	}

	err = UploadImage(ctx, storageClient, bucketName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Image uploaded successfully.")
}

func UploadImage(ctx context.Context, client *storage.Client, bucketName string, filePath string) error {
	// ファイルを読み込む
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("client.Bucket: %v", err)
	}
	objectName := "images/" + filePath
	object := bucket.Object(objectName)
	writer := object.NewWriter(ctx)
	defer writer.Close()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if _, err := io.Copy(writer, file); err != nil {
        return fmt.Errorf("io.Copy: %v", err)
    }

    return nil
}