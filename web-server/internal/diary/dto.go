package diary

type GetDiaryInput struct {
	Writer *string `form:"writer"`
	Week   *string `form:"week"`
}

type UnlockPasswordInput struct {
	Writer   string `json:"writer" binding:"required"`
	Week     string `json:"week" binding:"required"`
	Password string `json:"password" binding:"required"`
}
