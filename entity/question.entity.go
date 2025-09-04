package entity

type QuestionEntity struct {
	Id              uint
	Question        string
	PossibleAnswers []string
	CorrectAnswer   string
	Difficulty      string
	Categry         string
}
