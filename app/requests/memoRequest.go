package requests

// contetnフィールドは必須
type CreateMemoInput struct {
	Content string `json:"content" binding:"required"`
}

// updateの場合は必須ではない
type UpdateMemoInput struct {
	Content string `json:"content"`
}
