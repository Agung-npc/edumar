package repository

import (
	"database/sql"

	. "github.com/dwirobbin/edumar-backend/helper"
	. "github.com/dwirobbin/edumar-backend/model/domain"
)

type FeRepositoryImpl struct {
	DB *sql.DB
}

func NewFeRepository(db *sql.DB) *FeRepositoryImpl {
	return &FeRepositoryImpl{
		DB: db,
	}
}

func (repo *FeRepositoryImpl) FindCategories() ([]CategoryDomain, error) {
	var categories []CategoryDomain

	query := `SELECT id, name, description FROM categories;`
	rows, err := repo.DB.Query(query)
	PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var category CategoryDomain
		err := rows.Scan(&category.Id, &category.Name, &category.Description)
		PanicIfError(err)

		categories = append(categories, category)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *FeRepositoryImpl) FindCategoryById(categoryId uint) (CategoryDomain, error) {
	var category CategoryDomain

	query := `SELECT id, name, description FROM categories WHERE id = ?;`
	row := repo.DB.QueryRow(query, categoryId)
	err := row.Scan(&category.Id, &category.Name, &category.Description)
	PanicIfError(err)

	return category, nil
}

func (repo *FeRepositoryImpl) FindQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]QuizDomain, error) {
	query := `
	SELECT id, category_id, question, correct_answer FROM quizzes 
	WHERE category_id = ? ORDER BY id LIMIT ? OFFSET ?;`

	rows, err := repo.DB.Query(query, categoryId, limit, (page-1)*limit)
	PanicIfError(err)
	defer rows.Close()

	var quizzes []QuizDomain
	for rows.Next() {
		var quiz QuizDomain
		err := rows.Scan(&quiz.Id, &quiz.CategoryId, &quiz.Question, &quiz.CorrectAnswer)
		PanicIfError(err)

		quizzes = append(quizzes, quiz)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (repo *FeRepositoryImpl) FindIncorrectAnswersByQuizId(quizId uint) (IncorrectAnswerDomain, error) {
	query := `
	SELECT id, quiz_id, option_one, option_two FROM incorrect_answers 
	WHERE quiz_id = ?;`

	var incorrectAnswerDomain IncorrectAnswerDomain
	row := repo.DB.QueryRow(query, quizId)
	err := row.Scan(
		&incorrectAnswerDomain.Id, &incorrectAnswerDomain.QuizId,
		&incorrectAnswerDomain.OptionOne, &incorrectAnswerDomain.OptionTwo,
	)
	PanicIfError(err)

	return incorrectAnswerDomain, nil
}

func (repo *FeRepositoryImpl) SaveAnswerAttempt(userId uint, answersAttempt []AnswerAttemptDomain) (bool, error) {
	query := `DELETE FROM answer_attempts;`
	_, err := repo.DB.Exec(query)
	PanicIfError(err)

	query = `INSERT INTO answer_attempts (answer, quiz_id, user_id) VALUES (?, ?, ?);`

	for _, answerAttempt := range answersAttempt {
		_, err := repo.DB.Exec(query, answerAttempt.Answer, answerAttempt.QuizId, userId)
		PanicIfError(err)
	}

	return true, nil
}

func (repo *FeRepositoryImpl) SaveResult(duration string, userId, categoryId uint) (bool, error) {
	query := `
	INSERT INTO results (correct, wrong, duration, user_id, category_id)
	SELECT 
	(SELECT COUNT(aa.answer) FROM answer_attempts AS aa 
		INNER JOIN quizzes AS q 
			WHERE aa.answer = q.correct_answer) AS correct,
	(SELECT COUNT(aa.answer) FROM answer_attempts AS aa 
		INNER JOIN incorrect_answers AS ia
			WHERE aa.answer = ia.option_one OR aa.answer = ia.option_two) AS wrong,
	?, ?, ?;`

	_, err := repo.DB.Exec(query, duration, userId, categoryId)
	PanicIfError(err)

	return true, nil
}

func (repo *FeRepositoryImpl) FindResultByCategoryId(categoryId uint) (ResultDomain, error) {
	query := `SELECT * FROM results WHERE category_id = ? ORDER BY id DESC;`

	var result ResultDomain
	row := repo.DB.QueryRow(query, categoryId)
	err := row.Scan(&result.Id, &result.Correct, &result.Wrong, &result.Duration,
		&result.UserId, &result.CategoryId, &result.CreatedAt, &result.UpdatedAt,
	)
	PanicIfError(err)

	return result, nil
}

func (repo *FeRepositoryImpl) FindScoresBoardByCategoryId(categoryId uint) ([]ResultDomain, error) {
	query := `SELECT * FROM results WHERE category_id = ?;`

	rows, err := repo.DB.Query(query, categoryId)
	PanicIfError(err)
	defer rows.Close()

	var results []ResultDomain
	for rows.Next() {
		var result ResultDomain
		err := rows.Scan(
			&result.Id, &result.Correct, &result.Wrong, &result.Duration, &result.UserId,
			&result.CategoryId, &result.CreatedAt, &result.UpdatedAt,
		)
		PanicIfError(err)

		results = append(results, result)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
