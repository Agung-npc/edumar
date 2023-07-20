package seed

import (
	"database/sql"
	"log"

	. "github.com/dwirobbin/edumar-backend/helper"
)

func Seed(db *sql.DB) {
	queries := []string{
		`INSERT INTO categories (name, description)
		VALUES
		('VOCABULARY', "This course contains several questions that have been specially arranged to broaden your knowledge, especially about vocabulary."),
		('GRAMMAR', "This course contains several questions that have been specially arranged to broaden your knowledge, especially about grammar."),
		('TENSES', "This course contains several questions that have been specially arranged to broaden your knowledge, especially about tenses.");`,

		`INSERT INTO quizzes (category_id, question, correct_answer)
		VALUES
		(1, "Was Aminah  …  my house at this time yesterday ?", "visit"),
		(1, "Children always …….. their parents.", "are loving"),
		(1, "Oni Shahrial ... dubbing Shinchan voice at this time last year.", "was"),
		(1, "My daughter ... a lot of photographs of the Borobudur when she went there on her last vacation.", "took"),
		(1, "Were you watching strawberry shortcake cartoon when I came ?", "watching"),
		(1, "Diana's barbie is broken. Diana is very …… now.", "confuse"),
		(1, "The teacher's duty is to ….. the students in the school.", "work"),
		(1, "The carpet is …… . I want to clean it.", "dirty"),
		(1, "Kathy is a ….. . She teaches Math in our class. Every students love her.", "kind teacher"),
		(1, "I can't hear anything since my …… are sick.", "ears");`,

		`INSERT INTO incorrect_answers (quiz_id, option_one, option_two)
		VALUES
		(1, "visited", "visiting"),
		(2, "love", "loves"),
		(3, "is", "has been"),
		(4, "had taken", "has been taking"),
		(5, "has been watching", "watched"),
		(6, "sad", "charm"),
		(7, "play", "make"),
		(8, "large", "soft"),
		(9, "arrogant teacher", "emotional teacher"),
		(10, "nose", "eyes");`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		PanicIfError(err)
	}

	log.Println("Successfully seeded all table")
}
