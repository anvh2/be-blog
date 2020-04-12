package mysql

// UserData ...
type UserData struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
	DName    string `json:"dname"`
	Avatar   string `json:"avatar"`
	Role     int32  `jon:"role"`
}
