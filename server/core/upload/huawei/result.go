package huawei

import (
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	core "github.com/limeschool/easy-admin/server/core/upload"
)

func NewListObjectResult(output *obs.ListObjectsOutput) *core.ListObjectResult {
	return &core.ListObjectResult{
		Files:      getFiles(output.Contents),
		IsFinished: !output.IsTruncated,
	}
}

func getFiles(contents []obs.Content) []core.File {
	var files []core.File

	for _, content := range contents {
		if strings.HasSuffix(content.Key, "/") {
			continue
		}

		files = append(files, &File{content: content})
	}

	return files
}
