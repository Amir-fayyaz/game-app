package entity

type GameEntity struct {
	Id          uint
	CategoryId  uint
	QuestionsId []uint
	PlayersId   []uint
}
