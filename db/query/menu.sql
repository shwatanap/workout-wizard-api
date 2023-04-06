-- name: CreateMenu :exec
INSERT INTO menus (
	user_id,
	target,
	comment
) VALUES (
	?, ?, ?
);
