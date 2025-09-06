package entity

type AnswerEntity struct {
	Id         uint
	PlayerId   uint
	QuestionId uint
	choice     PossibleAnswerEntity
}
