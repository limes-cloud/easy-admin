package types

type CaptchaRequest struct {
	Name string `json:"name" binding:"required"`
}
