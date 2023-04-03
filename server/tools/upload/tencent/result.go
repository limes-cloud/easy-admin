package tencent

import (
	"strings"

	core "github.com/limeschool/easy-admin/server/tools/upload"
	cos "github.com/tencentyun/cos-go-sdk-v5"
)

func NewListObjectResult(r *cos.BucketGetResult) *core.ListObjectResult {
	return &core.ListObjectResult{
		Files:      getFiles(r.Contents),
		IsFinished: !r.IsTruncated,
	}
}

func getFiles(contents []cos.Object) []core.File {
	var files []core.File

	for _, content := range contents {
		if strings.HasSuffix(content.Key, "/") {
			continue
		}

		files = append(files, &File{object: content})
	}

	return files
}
