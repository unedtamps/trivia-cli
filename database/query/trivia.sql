-- name: FindTriviaWithSpecs :many
SELECT t.id, t.title, t.question , t.answer
FROM trivia t
WHERE t.dificulity_id = ? AND t.id NOT IN 
   (SELECT trivia_id FROM trivia_detail WHERE user_id = ? OR (trivia_status_id = 2 OR trivia_status_id = 3))
LIMIT ?;

-- name: FindChoicesOfTrivia :many
SELECT choice FROM trivia_choice WHERE trivia_id = ?;

-- name: CreateTrivia :execlastid
INSERT INTO trivia (title, question, answer, dificulity_id)
VALUES (?, ?, ?, ?);

-- name: CreateTriviaChoice :exec
INSERT INTO trivia_choice (trivia_id, choice)
VALUES (?, ?);

-- name: CreateTriviaDetail :exec
INSERT INTO trivia_detail (user_id, trivia_id, trivia_status_id)
VALUES (?, ?, ?);

-- name: UpdateTrivia :exec
UPDATE trivia SET title = ?, question = ?, answer = ?, dificulity_id = ? WHERE id = ?;

-- name: UpdateTriviaChoice :exec
UPDATE trivia_choice SET choice = ? WHERE id = ?;

-- name: FindAllSubmission :many
SELECT t.title, df.name as diffname, u.name FROM trivia t JOIN trivia_detail as td 
ON t.id = td.trivia_id JOIN user as u ON td.user_id = u.id JOIN dificulity as df ON df.id = t.dificulity_id
WHERE td.trivia_status_id = 2;

-- name: FindReviewSubmission :many
SELECT t.id, t.title, t.question, t.answer, t.dificulity_id, td.trivia_status_id, tc.id as choice_id, tc.choice FROM trivia t 
JOIN trivia_detail as td ON t.id = td.trivia_id JOIN trivia_choice as tc ON t.id = tc.trivia_id
WHERE td.trivia_status_id = 2;

-- name: UpdateTrivaStatus :exec
UPDATE trivia_detail SET trivia_status_id = ? WHERE trivia_id = ?;

