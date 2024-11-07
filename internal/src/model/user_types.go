package model

// UserData stores all user data
type UserData struct {
	UserName string `json:"username"`
	UserID   int64  `json:"user_id"`
	Pwd      string `json:"pwd"`
}
