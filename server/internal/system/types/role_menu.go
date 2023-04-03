package types

type AddRoleMenuRequest struct {
	RoleID  int64   `json:"role_id" binding:"required"`
	MenuIds []int64 `json:"menu_ids" binding:"required"`
}

type RoleMenuIdsRequest struct {
	RoleID int64 `json:"role_id" form:"role_id" binding:"required"`
}

type RoleMenuRequest struct {
	RoleID int64 `json:"role_id" form:"role_id" binding:"required"`
}
