package response

type CreateMenuResponse struct {
	ID       int32                   `json:"id"`
	UserID   int32                   `json:"user_id"`
	Workouts []CreateWorkoutResponse `json:"workouts"`
	Target   string                  `json:"target"`
	Comment  string                  `json:"comment"`
}
