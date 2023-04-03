package types

type UploadRequest struct {
	Dir string `json:"dir" form:"dir"  binding:"required"`
}
