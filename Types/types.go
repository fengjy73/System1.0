package Types

type UserAuthority int

const (
	// Owner For example,if a user is an Owner,his UserAuthority&1 = 1
	Owner         UserAuthority = 1
	Administrator UserAuthority = 2
	Registered    UserAuthority = 4
	Tourist       UserAuthority = 8
	// A user can be an Administrator and Registered in the meantime
)

type ErrNo int

const (
	OK                      ErrNo = 0
	ParamInvalid            ErrNo = 1
	DatabaseConnectionErrNo ErrNo = 2
	PasswordWrong           ErrNo = 3
	PermissionDenied        ErrNo = 4

	UnknownErrNo ErrNo = -1
)

type User struct {
	UserID   string
	UserName string
	NickName string
	Password string
}
type LoginRequest struct {
	UserName string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

// LoginResponse after get login in successfully, return Cookie Session
type LoginResponse struct {
	Code ErrNo
	Data struct {
		UserID string
	}
}
type LogoutRequest struct {
}
type LogoutResponse struct {
	Code ErrNo
}

// WhoAmIRequest use for test and retrieve Password
type WhoAmIRequest struct {
}
type WhoAmIResponse struct {
	Code ErrNo
	Data User
}

type CreateUserRequest struct {
	NickName      string
	Username      string
	Password      string
	UserAuthority UserAuthority
}
type CreateUserResponse struct {
	Code ErrNo
	Data struct {
		UserID string
	}
}
type GetUserRequest struct {
	UserID string
}
type GetUserResponse struct {
	Code ErrNo
	Data User
}
type GetUserListRequest struct {
	Offset int
	Limit  int
}
type GetUserListResponse struct {
	Code ErrNo
	Data struct {
		UserList []User
	}
}
type UpdateUserRequest struct {
	UserID   string
	NickName string
}
type UpdateUserResponse struct {
	Code ErrNo
}
type DeleteUserRequest struct {
	UserID string
}
type DeleteUserResponse struct {
	Code ErrNo
}
