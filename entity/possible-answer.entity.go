package entity

import (
	types "main/types"
)

type PossibleAnswerEntity struct {
	Id     uint
	Text   string
	choice types.PossibleAnswerChoiceType
}
