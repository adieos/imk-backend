package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL            = "failed verify email"
	MESSAGE_FAILED_RESET_PASSWORD          = "failed reset password"
	MESSAGE_FAILED_EMAIL_NOT_FOUND         = "email not found"
	MESSAGE_FAILED_TOKEN_EXPIRED           = "token expired"
	MESSAGE_FAILED_FORGET_PASSWORD         = "failed handle forget password"
	MESSAGE_FAILED_TOKEN_NOT_ALLOWED       = "token not allowed"
	MESSAGE_FAILED_GET_MOODLE_CREDS        = "failed get moodle credentials"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
	MESSAGE_SUCCESS_RESET_PASSWORD          = "success reset password"
	MESSAGE_SUCCESS_FORGET_PASSWORD         = "success handle forget password"
	MESSAGE_SUCCESS_GET_MOODLE_CREDS        = "success get moodle credentials"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotAdmin           = errors.New("user not admin")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrHashPasswordFailed     = errors.New("failed to hash password")
	ErrParseUUID              = errors.New("error parsing uuid")
	ErrUserIdEmpty            = errors.New("user id empty")
	ErrRoleNotAllowed         = errors.New("role not allowed: ")
	ErrGetMoodleCreds         = errors.New("failed get moodle credentials")
)

type (
	UserCreateRequest struct {
		Name     string `json:"name" form:"name" binding:"required,max=700"`
		Email    string `json:"email" form:"email" binding:"required,max=100"`
		Password string `json:"password" form:"password" binding:"required,max=100"`
	}

	UserResponse struct {
		Name  string `json:"name" form:"name" binding:"required"`
		Email string `json:"email" form:"email" binding:"required"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required,max=100"`
		Password string `json:"password" form:"password" binding:"required,max=100"`
	}

	UserLoginResponse struct {
		Token string `json:"token" form:"token" binding:"required,max=100"`
	}
)
