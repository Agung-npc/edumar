package domain

import "time"

type UserDomain struct {
	Id        uint
	Username  string
	Email     string
	Password  string
	Loggedin  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryDomain struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QuizDomain struct {
	Id            uint
	CategoryId    uint
	Question      string
	CorrectAnswer string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type IncorrectAnswerDomain struct {
	Id        uint
	QuizId    uint
	OptionOne string
	OptionTwo string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AnswerAttemptDomain struct {
	Id        uint
	Answer    string
	QuizId    uint
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResultDomain struct {
	Id         uint
	UserId     uint
	CategoryId uint
	Correct    uint
	Wrong      uint
	Duration   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
