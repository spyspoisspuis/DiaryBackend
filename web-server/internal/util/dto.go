package util

type DiaryStruct struct {
	Writer   string `json:"writer" binding:"required"`
	Week     string `json:"week" binding:"required"`
	Header   string `json:"header" `
	Context  string `json:"context" binding:"required"`
	Footer   string `json:"footer"`
	Status   string `json:"status"`
}
