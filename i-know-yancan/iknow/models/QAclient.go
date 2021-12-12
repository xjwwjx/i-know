package models

type QuestionHot struct {
	Great    int
	UserName string
	QUE      string
	AnsNum   int
}

type Like struct {
	QuestionID int `form:"questionid" json:"questionid" binding:"required"`
}

type Search struct {
	KeyWords string `form:"keywords" json:"keywords" binding:"required"`
}
type QuestionsParam struct {
	QUE      string `form:"que" json:"que" binding:"required"`
	UserName string `form:"username" json:"username" binding:"required"`
}
type DelQuestionsParam struct {
//QUE      string `form:"que" json:"que" binding:"required"`
ID       uint   `form:"id" json:"id" binding:"required"`
//UserName string `form:"username" json:"username" binding:"required"`
}
type AnswersParam struct {
	ANS        string `form:"ans" json:"ans" binding:"required"`
	UserName   string `form:"userName" json:"userName" binding:"required"`
	QuestionID int
}
type DelAnswersParam struct {
	//ANS        string `form:"ans" json:"ans" binding:"required"`
	//UserName   string `form:"userName" json:"userName" binding:"required"`
	//QuestionID int
	ID       uint   `form:"id" json:"id" binding:"required"`
}
type CollectParam struct {
	UserID     int `form:"userID" json:"userid" binding:"required"`
	QuestionID int `form:"questionID" json:"questionid" binding:"required"`
}
