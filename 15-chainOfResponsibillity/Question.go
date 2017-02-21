package main

type Question struct {
	QuestionContents string
	Level            int
}

func NewQuestion(questionContents string, level int) *Question {
	return &Question{QuestionContents: questionContents, Level: level}
}
