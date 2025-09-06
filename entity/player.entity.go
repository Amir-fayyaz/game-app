package entity

type Player struct {
	Id      uint
	UserId  uint
	GameId  uint
	Score   uint
	Answers []AnswerEntity
}
