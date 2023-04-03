package minio

import (
	"context"

	core "github.com/limeschool/easy-admin/server/tools/upload"
	"github.com/minio/minio-go/v7"
)

const maxKeys = 1000

type Chunks struct {
	bucket string
	prefix string

	client *minio.Client
}

func NewChunks(bucket string, prefix string, client *minio.Client) core.Chunks {
	return &Chunks{
		bucket: bucket,
		prefix: prefix,
		client: client,
	}
}

func (c *Chunks) Chunk() (*core.ListObjectResult, error) {
	opts := minio.ListObjectsOptions{
		Prefix:  c.prefix,
		MaxKeys: maxKeys,
	}

	ch := c.client.ListObjects(context.Background(), c.bucket, opts)

	var result []minio.ObjectInfo
	for {
		v, ok := <-ch
		if !ok {
			break
		}

		result = append(result, v)
	}

	return NewListObjectResult(result, false), nil
}
