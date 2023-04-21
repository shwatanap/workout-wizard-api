package response

type CreateWorkoutResponse struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
