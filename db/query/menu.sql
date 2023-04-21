-- name: CreateMenu :exec
INSERT INTO menus (
	user_id,
	target,
	comment
) VALUES (
	?, ?, ?
);

-- name: GetCreatedMenuID :one
SELECT LAST_INSERT_ID();

-- name: GetMenu :one
SELECT (
	id,
	user_id,
	target,
	comment
)
FROM menus
WHERE id = ?;
