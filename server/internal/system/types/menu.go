package types

type AddMenuRequest struct {
	ParentID   int64  `json:"parent_id"  binding:"required"`
	Title      string `json:"title"  binding:"required"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Type       string `json:"type"  binding:"required"`
	Permission string `json:"permission"`
	Method     string `json:"method"`
	Component  string `json:"component"`
	Redirect   string `json:"redirect"`
	Weight     int    `json:"weight"`
	IsHidden   bool   `json:"is_hidden"`
	IsCache    bool   `json:"is_cache"`
	IsHome     bool   `json:"is_home"`
}

type UpdateMenuRequest struct {
	ID         int64   `json:"id" binding:"required"`
	ParentID   int64   `json:"parent_id"  binding:"required"`
	Title      string  `json:"title"  binding:"required"`
	Icon       string  `json:"icon"`
	Path       string  `json:"path"`
	Name       string  `json:"name"`
	Type       string  `json:"type"  binding:"required"`
	Permission string  `json:"permission"`
	Method     string  `json:"method"`
	Component  string  `json:"component"`
	Redirect   *string `json:"redirect"`
	Weight     *int    `json:"weight"`
	IsHidden   *bool   `json:"is_hidden"`
	IsCache    *bool   `json:"is_cache"`
	IsHome     *bool   `json:"is_home"`
}

type DeleteMenuRequest struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}
