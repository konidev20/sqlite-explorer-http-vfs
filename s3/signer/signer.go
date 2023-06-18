package main

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type Client struct {
	*minio.Client
}

func New(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*Client, error) {
	// Initialize the client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("\nFailed to initialize MINIO client: %w", err)
		return nil, errors.Wrapf(err, "Failed to initialize MINIO client")
	}

	return &Client{client}, nil
}

func (c *Client) GetPresignedURL(ctx context.Context, bucketname, objectname string, ttl time.Duration) (url string, err error) {

	// Generate the pre-signed URL
	presignedURL, err := c.PresignedGetObject(ctx, bucketname, objectname, ttl, nil)
	if err != nil {
		log.Fatalf("Failed to generate pre-signed URL:%w", err)
		return "", errors.Wrapf(err, "Failed to generate pre-signed URL")
	}

	return presignedURL.String(), nil
}
