package migrate

import (
	"database/sql"
	"log"

	. "github.com/dwirobbin/edumar-backend/helper"
)

func Migrate(db *sql.DB) {
	queries := []string{
		`SET FOREIGN_KEY_CHECKS=0;`,
		`CREATE TABLE IF NOT EXISTS users (
				id INT NOT NULL AUTO_INCREMENT,
				username VARCHAR(100) NOT NULL,
				email VARCHAR(100) NOT NULL UNIQUE,
				password TEXT NOT NULL,
				loggedin BOOLEAN NOT NULL DEFAULT FALSE,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE categories (
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(100) NOT NULL UNIQUE,
				description TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS quizzes (
				id INT NOT NULL AUTO_INCREMENT,
				category_id INT NOT NULL,
				question TEXT NOT NULL,
				correct_answer TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS incorrect_answers (
				id INT NOT NULL AUTO_INCREMENT,
				quiz_id INT NOT NULL,
				option_one TEXT NOT NULL,
				option_two TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS answer_attempts (
				id INT NOT NULL AUTO_INCREMENT,
				answer TEXT NOT NULL,
				quiz_id INT NOT NULL,
				user_id INT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
		`CREATE TABLE IF NOT EXISTS results (
				id INT NOT NULL AUTO_INCREMENT,
			  correct INT NOT NULL,
				wrong INT NOT NULL,
				duration VARCHAR(20) NOT NULL,
				user_id INT NOT NULL,
				category_id INT NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY (id)
			);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		PanicIfError(err)
	}

	log.Println("Successfully migrated all table")
}
