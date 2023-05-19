package types

type AllFileRequest struct {
	Name     string `json:"name" form:"name" sql:"like '%?%'"`
	ParentID int64  `json:"parent_id" form:"parent_id"`
}

type AddFileRequest struct {
	Src      string `json:"src" binding:"required"`
	Name     string `json:"name" binding:"required"`
	IsDir    *bool  `json:"is_dir" binding:"required"`
	ParentID int64  `json:"parent_id" binding:"required"`
	Size     int64  `json:"size"`
	Suffix   string `json:"suffix"`
	Mime     string `json:"mime"`
}

type UpdateFileRequest struct {
	ID       int64  `json:"id"  binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parent_id" binding:"required"`
}

type DeleteFileRequest struct {
	ID int64 `json:"id"`
}
