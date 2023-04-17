package types

type PageNoticeRequest struct {
	Page     int    `json:"page" form:"page" binding:"required" sql:"-"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required,max=50"  sql:"-"`
	Title    string `json:"title" form:"title" sql:"like '%?%'"`
	Status   *bool  `json:"status" form:"status"`
	IsRead   *bool  `json:"is_read" form:"is_read" sql:"-"`
	Start    int64  `json:"start" form:"start" sql:"> ?" column:"created_at"`
	End      int64  `json:"end" form:"end" sql:"< ?" column:"created_at"`
}

type GetNoticeRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type AddNoticeRequest struct {
	Type    string `json:"type"  binding:"required"`
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content"  binding:"required"`
}

type UpdateNoticeRequest struct {
	ID      int64  `json:"id"  binding:"required"`
	Type    string `json:"type"`
	Status  *bool  `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteNoticeRequest struct {
	ID int64 `json:"id"`
}
