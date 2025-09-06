package entity

import "time"

type GameEntity struct {
	Id          uint
	CategoryId  uint
	QuestionsId []uint
	PlayersId   []uint
	StartTime   time.Time
}
