package types

type PossibleAnswerChoiceType uint8

func (p PossibleAnswerChoiceType) isValid() bool {
	return (p >= 1 && p <= 4)
}
