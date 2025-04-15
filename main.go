package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

/*
storageAccountUrl       = "https://sad8ayv7a7test1.blob.core.windows.net/"
containerName = "ct-hosting-cts-d8ay-v7a7-test1"
blobName      = "test-up.txt"
localFile    = "./test-up.txt"
*/
func main() {
	storageAccountUrl := flag.String("sa", "", "The Storage Account URL")
	containerName := flag.String("container", "", "The container Name")
	blobName := flag.String("remote", "", "The Blob name to upload")
	localFile := flag.String("local", "", "The local file to upload")
	flag.Parse()

	// authenticate with Azure Active Directory
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// create a client for the specified storage account
	client, err := azblob.NewClient(*storageAccountUrl, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(*localFile, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// upload the file to the specified container with the specified blob name
	_, err = client.UploadFile(context.TODO(), *containerName, *blobName, file, nil)
	if err != nil {
		log.Fatal(err)
	}
}
