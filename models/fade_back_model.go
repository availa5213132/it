package models

type FadeBackModel struct {
	MODEL
	Email        string `json:"email"`
	Content      string `json:"content"`
	ApplyContent string `json:"apply_content"` // 回复的内容
	IsApply      bool   `json:"is_apply"`      // 是否回复
}
