package errors

// Error ...
type Error int32

// Define ...
const (
	Success = 1

	EmptyUsername       = -100
	EmptyPassword       = -101
	EmptyDName          = -102
	EmptyBlogHeader     = -103
	EmptyBlogSubtitle   = -104
	EmptyBlogBackground = -105
	EmptyBlogContent    = -106
	EmptyBlogID         = -107

	FailedLogin        = -200
	FailedRegisterUser = -201
	FailedGenUserID    = -202
	FailedGenBlogID    = -203
	FailedCreateBlog   = -204
	FailedGetBlog      = -205

	InvalidUsername = -300
	InvalidPassword = -301
	InvalidReadTime = -302

	ErrorSignToken        = -400
	ErrorPasswordNotMatch = -401
)

var messageVN = map[int32]string{
	Success: "Thành công",

	EmptyUsername:       "Tên đăng nhập không hợp lệ",
	EmptyPassword:       "Mật khẩu không hợp lệ",
	EmptyDName:          "Họ và tên không hợp lệ",
	EmptyBlogHeader:     "Blog header không hợp lệ",
	EmptyBlogSubtitle:   "Blog subtitle không hợp lệ",
	EmptyBlogBackground: "Blog background không hợp lệ",
	EmptyBlogContent:    "Blog content không hợp lệ",
	EmptyBlogID:         "BlogID không hợp lệ",

	FailedLogin:        "Đăng nhập thất bại",
	FailedRegisterUser: "Đăng kí thất bại",
	FailedGenUserID:    "Đăng kí thất bại",
	FailedGenBlogID:    "Đăng blog thất bại",
	FailedCreateBlog:   "Đăng blog thất bại",
	FailedGetBlog:      "Blog không tồn tại",

	InvalidUsername: "",
	InvalidPassword: "Mật khẩu không hợp lệ",
	InvalidReadTime: "",

	ErrorSignToken:        "Đăng nhập thất bại",
	ErrorPasswordNotMatch: "Mật khẩu không khớp",
}

var messageEN = map[int32]string{
	EmptyUsername: "Invalid username",
	EmptyPassword: "Invalid password",
}

// GetMessage ...
func GetMessage(code int32) string {
	return messageVN[code]
}
