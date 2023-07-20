package web

type LoginResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type QuizResponse struct {
	Id               uint     `json:"id"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type AnswerAttemptResponse struct {
	Correct  uint   `json:"correct"`
	Wrong    uint   `json:"wrong"`
	Duration string `json:"duration"`
}

type ScoreBoardResponse struct {
	Username string `json:"username"`
	Score    uint   `json:"score"`
	Duration string `json:"duration"`
}
