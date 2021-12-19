package models

type QuestionHot struct {
	Great    int
	UserName string
	QUE      string
	AnsNum   int
}

type Like struct {
	QuestionID int
}

type Search struct {
	KeyWords string `form:"keywords" json:"keywords" binding:"required"`
}
type QuestionsParam struct {
	QUE      string `form:"que" json:"que" binding:"required"`
	UserName string `form:"username" json:"username" binding:"required"`
}

type AnswersParam struct {
	ANS        string `form:"ans" json:"ans" binding:"required"`
	UserName   string `form:"userName" json:"userName" binding:"required"`
	QuestionID int
}
type CollectParam struct {
	UserID     int `form:"userID" json:"userID" binding:"required"`
	QuestionID int `form:"questionID" json:"questionID" binding:"required"`
}
