-- name: CreateMessage :one
INSERT INTO message (thread_id, sender, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateThread :one
INSERT INTO thread (topic)
VALUES ($1)
RETURNING *;

-- name: GetMessageByID :one
SELECT * FROM message
WHERE id = $1;

-- name: GetMessagesByThread :many
SELECT * FROM message
WHERE thread_id = $1
ORDER BY created_at DESC;

-- name: DeleteMessage :exec
DELETE FROM message WHERE id = $1;

-- name: UpdateMessage :exec
UPDATE message 
SET content = $2
WHERE id = $1
RETURNING *;

-- name: GetThreadById :one
SELECT * FROM thread
WHERE id = $1;

-- name: DeleteAll :exec
DELETE FROM message;    