package response

// User
type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    uint   `json:"age"`
}
