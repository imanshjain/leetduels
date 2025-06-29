package models

type Question struct {
	QuestionId   string
	Difficulty   string
	Description  string            // includes the examples
	TestCases    map[string]string // For now keeping a string input and output (probbaly parse them when testing)
	TimeLimit    int               // in minutes
	AvgTimeTaken *int              // optional
}

func GetQuestion() *Question {
	return &Question{
		QuestionId:   "",
		Difficulty:   "",
		Description:  "",
		TestCases:    make(map[string]string),
		TimeLimit:    0,
		AvgTimeTaken: nil,
	}
}
