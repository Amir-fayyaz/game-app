package types

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (q QuestionDifficulty) isValid() bool {
	return (q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard)
}
