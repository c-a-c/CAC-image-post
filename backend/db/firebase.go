package db

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// Firebaseとの接続
func NewFirebase() (*storage.Client, string, error) {
	// 環境変数の読み込み
	if os.Getenv(("GO_ENV")) == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Firebaseの設定
	ctx := context.Background()
	credentialsFile := os.Getenv("FIREBASE_ADMIN_SDK")
    sa := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// Storageのクライアントを取得
	storageClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")
	return storageClient, bucketName, nil
}