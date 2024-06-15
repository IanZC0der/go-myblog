package user

// User represents the user entity
type User struct {
	Id        int64 `json:"id"`
	CreatedAt int64 `json:"created_at"` //not datetime as it increases complexity (timezones, etc.)
	UpdateAt  int64 `json:"updated_at"`
	*CreateUserRequest
}
