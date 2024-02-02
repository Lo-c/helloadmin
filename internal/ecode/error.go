package ecode

//go:generate stringer -type ErrCode -linecomment

type ErrCode int

func (e ErrCode) Int() int {
	return int(e)
}

func (e ErrCode) Code() ErrCode {
	return e
}

func (e ErrCode) Error() string {
	return e.String()
}

const (
	ErrSuccess             ErrCode = 0   //Success
	ErrBadRequest          ErrCode = 400 //Bad Request
	ErrUnauthorized        ErrCode = 401 //Unauthorized
	ErrNotFound            ErrCode = 404 //Not Found
	ErrInternalServerError ErrCode = 500 //Internal Server Error

	ErrEmailAlreadyUse   ErrCode = iota + 10001 //The email is already in use
	ErrPasswordIncorrect                        //The password is incorrect
	ErrUserNotFound                             //The user does not exist
	ErrRoleHasUser                              //The role has users and cannot be deleted
)
