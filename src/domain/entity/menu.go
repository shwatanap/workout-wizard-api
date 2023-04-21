package entity

type Menu struct {
	ID       int32      `json:"id"`
	UserID   int32      `json:"user_id"`
	Workouts []*Workout `json:"workouts"`
	Target   string     `json:"target"`
	Comment  string     `json:"comment"`
}
