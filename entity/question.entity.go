package entity

import (
	types "main/types"
)

type QuestionEntity struct {
	Id              uint
	Text            string
	PossibleAnswers []PossibleAnswerEntity
	CorrectAnswerId uint
	Difficulty      types.QuestionDifficulty
	CategryId       uint
}
