package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"sort"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Utils struct {
	bucketClient *storage.Client
}

const (
	bucket   = "crabstash-staging"
	endpoint = "https://storage.googleapis.com"
	region   = "europe-central2"
)

type SortingStruct struct {
	Generation int64
	Created    time.Time
}

func InitBucket() Utils {
	ctx := context.Background()
	sess, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Error while connecting to bucket: %e", err)
	}

	return Utils{sess}
}

func (utils *Utils) UploadFile(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	file, err := fileHeader.Open()

	if err != nil {
		return "", fmt.Errorf("error while opening file: %s", err)
	}

	o := utils.bucketClient.Bucket(bucket).Object(fileName)
	ctx := context.Background()
	wc := o.NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("error while copying file to buffer: %s", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("error while saving file to bucket: %s", err)
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, fileName), nil

}

func (utils *Utils) DeleteFile(fileName string) error {
	ctx := context.Background()
	o := utils.bucketClient.Bucket(bucket).Object(fileName)

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("error while deleting file: %e", err)
	}

	return nil

}

func (utils *Utils) RestoreFile(fileName string) error {
	ctx := context.Background()

	it := utils.bucketClient.Bucket(bucket).Objects(ctx, &storage.Query{
		Versions: true,
		Prefix:   fileName,
	})

	var sorting []SortingStruct

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return fmt.Errorf("error while iterating versions: %e", err)
		}

		sorting = append(sorting, SortingStruct{
			Generation: attrs.Generation,
			Created:    attrs.Created,
		})
	}

	sort.Slice(sorting, func(i, j int) bool {
		return sorting[i].Created.After(sorting[j].Created)
	})

	src := utils.bucketClient.Bucket(bucket).Object(fileName)

	if _, err := src.CopierFrom(src.Generation(sorting[1].Generation)).Run(ctx); err != nil {
		return fmt.Errorf("error while restoring version: %e", err)
	}

	if err := src.Generation(sorting[0].Generation).Delete(ctx); err != nil {
		return fmt.Errorf("error while deleting file: %e", err)
	}

	return nil
}
