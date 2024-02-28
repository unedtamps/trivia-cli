-- name: GetUserById :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserTriviaDetails :one
SELECT u.name, t.title FROM user u 
JOIN trivia_detail ON u.id = ut.user_id 
JOIN trivia t ON ut.trivia_id = t.id 
WHERE u.id = ?;

-- name: TriviaByDificulity :many
SELECT * FROM trivia WHERE dificulity_id = ?;

-- name: SelectUserByEmail :one
SELECT * FROM user WHERE email = ?;

-- name: CreateUser :execlastid
INSERT INTO user (name, user_name, role_id, email, password) 
VALUES (?,?,?,?,?);

-- name: MigrateAdmin :execlastid
INSERT INTO user (name, user_name, role_id, email, password) VALUES (?, ?, ?, ?, ?);

-- name: CreateUserDetail :exec
INSERT INTO user_detail (user_id) VALUES (?);

-- name: CheckIfAdmin :one
SELECT role_id FROM user WHERE id = ?;

-- name: FindUserDetailById :one
SELECT trivia_round,right_answer, wrong_answer FROM user_detail WHERE user_id = ?;

-- name: UpdateUserDetail :exec
UPDATE user_detail SET trivia_round = ?, right_answer = ?, wrong_answer = ? WHERE user_id = ?;

-- name: SelectUserRankDesc :many
SELECT u.name, ud.trivia_round, ud.right_answer , (ud.wrong_answer + ud.right_answer) as total_answer
FROM user u JOIN user_detail ud ON u.id = ud.user_id
WHERE ud.trivia_round > 0 ORDER BY ud.right_answer DESC, ud.trivia_round DESC;

-- name: SelectUserRankAsc :many
SELECT u.name, ud.trivia_round, ud.right_answer, (ud.wrong_answer + ud.right_answer) as total_answer
FROM user u JOIN user_detail ud ON u.id = ud.user_id
WHERE ud.trivia_round > 0 ORDER BY ud.right_answer ASC, ud.trivia_round ASC;

-- name: SelectUserAboveAverage :many
SELECT u.name, ud.trivia_round, ud.right_answer, ud.wrong_answer
FROM user u JOIN user_detail ud ON u.id = ud.user_id
WHERE ud.trivia_round > 0 AND ud.right_answer/(ud.right_answer + ud.wrong_answer) >= (
  SELECT AVG(right_answer/(right_answer+wrong_answer)) FROM user_detail WHERE trivia_round > 0
);

-- name: SelectUserBelowAverage :many
SELECT u.name, ud.trivia_round, ud.right_answer, ud.wrong_answer
FROM user u JOIN user_detail ud ON u.id = ud.user_id
WHERE ud.trivia_round > 0 AND ud.right_answer/(ud.right_answer + ud.wrong_answer) < (
  SELECT AVG(right_answer/(right_answer+wrong_answer)) FROM user_detail WHERE trivia_round > 0
);


-- name: FindMyTrivia :many
SELECT t.id, t.title, t.dificulity_id, td.trivia_status_id FROM trivia t
JOIN trivia_detail td ON t.id = td.trivia_id
WHERE td.user_id = ? AND (td.trivia_status_id = 2 OR td.trivia_status_id = 3);

-- name: DeleteTrivia :exec
DELETE FROM trivia WHERE id = ?;

-- name: DeleteTriviaChoice :exec
DELETE FROM trivia_choice WHERE trivia_id = ?;

-- name: DeleteTriviaDetail :exec
DELETE FROM trivia_detail WHERE trivia_id = ?;
