package apperrors

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrPasswordWrong = errors.New("password not correct")
	ErrAuthorization = errors.New("authorization error")
	ErrCheckWord     = errors.New("spelling error in the message")
)
