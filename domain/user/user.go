package user

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}