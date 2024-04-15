package dto

type EditorAccountDTO struct {
	Id       *int64  `json:"id"`
	Nickname *string `json:"nickname" binding:"required" message:"昵称不能为空"`
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
