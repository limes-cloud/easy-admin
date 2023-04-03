package types

type PageUserRequest struct {
	Page     int    `json:"page" form:"page" binding:"required" sql:"-"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required,max=50"  sql:"-"`
	TeamID   int64  `json:"team_id" form:"team_id"`
	RoleID   int64  `json:"role_id" form:"role_id"`
	Name     string `json:"name" form:"name" sql:"like '%?%'"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Sex      *bool  `json:"sex" form:"sex"`
	Status   *bool  `json:"status" form:"status"`
	Start    int64  `json:"start" form:"start" sql:"> ?" column:"created_at"`
	End      int64  `json:"end" form:"end" sql:"< ?" column:"created_at"`
}

type GetUserRequest struct {
	ID int64 `json:"id" form:"id"  binding:"required"`
}

type AddUserRequest struct {
	TeamID   int64  `json:"team_id" binding:"required"`
	RoleID   int64  `json:"role_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sex      *bool  `json:"sex" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"  binding:"required"`
	Status   *bool  `json:"status"  binding:"required"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"id"  binding:"required"`
	TeamID   int64  `json:"team_id"`
	RoleID   int64  `json:"role_id"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Sex      *bool  `json:"sex"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   *bool  `json:"status"`
	Password string `json:"password"`
}

type UpdateUserinfoRequest struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Sex      *bool  `json:"sex"`
}

type DeleteUserRequest struct {
	ID int64 `json:"id"`
}

type UserLoginRequest struct {
	Phone     string `json:"phone"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

type UpdatePasswordRequest struct {
	Password  string `json:"password"`
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}
