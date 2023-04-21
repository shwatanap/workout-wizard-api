-- name: CreateWorkout :exec
INSERT INTO workouts (
	menu_id,
	name,
	detail
) VALUES (
	?, ?, ?
);

-- name: GetCreatedWorkoutID :one
SELECT LAST_INSERT_ID();
