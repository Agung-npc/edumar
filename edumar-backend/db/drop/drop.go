package drop

import (
	"database/sql"
	"log"

	. "github.com/dwirobbin/edumar-backend/helper"
)

func Drop(db *sql.DB) {
	queries := []string{
		`SET FOREIGN_KEY_CHECKS = 0;`,
		`DROP TABLE IF EXISTS users;`,
		`DROP TABLE IF EXISTS categories;`,
		`DROP TABLE IF EXISTS quizzes;`,
		`DROP TABLE IF EXISTS incorrect_answers;`,
		`DROP TABLE IF EXISTS answer_attempts;`,
		`DROP TABLE IF EXISTS results;`,
		`SET FOREIGN_KEY_CHECKS = 1;`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		PanicIfError(err)
	}

	log.Println("Successfully dropped all table")
}
