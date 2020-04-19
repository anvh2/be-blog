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
	EmptyEmail          = -108

	FailedLogin         = -200
	FailedRegisterUser  = -201
	FailedGenUserID     = -202
	FailedGenBlogID     = -203
	FailedCreateBlog    = -204
	FailedViewBlog      = -205
	FailedHashPassword  = -206
	FailedGetBlog       = -207
	FailedGetNumOfBlogs = -208

	InvalidUsername    = -300
	InvalidPassword    = -301
	InvalidReadTime    = -302
	InvalidOffsetLimit = -303

	ErrorSignToken        = -400
	ErrorPasswordNotMatch = -401
	ErrorPermissionDenied = -402
)

var message = map[int32]string{
	Success: "OK",

	EmptyUsername:       "EMPTY_USERNAME",
	EmptyPassword:       "EMPTY_PASSWORD",
	EmptyDName:          "EMPTY_DISPLAY_NAME",
	EmptyBlogHeader:     "EMPTY_BLOG_HEADER",
	EmptyBlogSubtitle:   "EMPTY_BLOG_SUBTITLE",
	EmptyBlogBackground: "EMPTY_BLOG_BACKGROUND",
	EmptyBlogContent:    "EMPTY_BLOG_CONTENT",
	EmptyBlogID:         "EMPTY_BLOG_ID",
	EmptyEmail:          "EMPTY_EMAIL",

	FailedLogin:         "FAILED_LOGIN",
	FailedRegisterUser:  "FAILED_REGISTER",
	FailedGenUserID:     "FAILED_GEN_USER_ID",
	FailedGenBlogID:     "FAILED_GEN_BLOG_ID",
	FailedCreateBlog:    "FAILED_CREATE_BLOG",
	FailedViewBlog:      "FAILED_VIEW_BLOG",
	FailedHashPassword:  "FAILED_HASH_PASSWORD",
	FailedGetBlog:       "FAILED_GET_BLOG",
	FailedGetNumOfBlogs: "FAILED_GET_NUM_OF_BLOGS",

	InvalidUsername:    "INVALID_USERNAME",
	InvalidPassword:    "INVALID_PASSWORD",
	InvalidReadTime:    "INVALID_READ_TIME",
	InvalidOffsetLimit: "INVALID_OFFSET_LIMIT",

	ErrorSignToken:        "ERROR_SIGN_TOKEN",
	ErrorPasswordNotMatch: "ERROR_PASSWORD_NOT_MATCH",
	ErrorPermissionDenied: "ERROR_PERMISSION_DENIED",
}

var detail = map[int32]string{
	Success: "Thành công",

	EmptyUsername:       "Tên đăng nhập không hợp lệ",
	EmptyPassword:       "Mật khẩu không hợp lệ",
	EmptyDName:          "Họ và tên không hợp lệ",
	EmptyBlogHeader:     "Blog header không hợp lệ",
	EmptyBlogSubtitle:   "Blog subtitle không hợp lệ",
	EmptyBlogBackground: "Blog background không hợp lệ",
	EmptyBlogContent:    "Blog content không hợp lệ",
	EmptyBlogID:         "BlogID không hợp lệ",
	EmptyEmail:          "",

	FailedLogin:         "Đăng nhập thất bại",
	FailedRegisterUser:  "Đăng kí thất bại",
	FailedGenUserID:     "Đăng kí thất bại",
	FailedGenBlogID:     "Đăng blog thất bại",
	FailedCreateBlog:    "Đăng blog thất bại",
	FailedViewBlog:      "Blog không tồn tại",
	FailedGetBlog:       "",
	FailedGetNumOfBlogs: "",

	InvalidUsername:    "",
	InvalidPassword:    "Mật khẩu không hợp lệ",
	InvalidReadTime:    "",
	InvalidOffsetLimit: "",

	ErrorSignToken:        "Đăng nhập thất bại",
	ErrorPasswordNotMatch: "Mật khẩu không khớp",
	ErrorPermissionDenied: "",
}

// GetMessage ...
func GetMessage(code int32) string {
	return message[code]
}

// GetDetail ...
func GetDetail(code int32) string {
	return detail[code]
}
