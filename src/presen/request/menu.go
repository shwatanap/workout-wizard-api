package request

type CreateMenuRequest struct {
	UserID  int32  `json:"user_id"`
	Target  string `json:"target"`
	Comment string `json:"comment"`
}
