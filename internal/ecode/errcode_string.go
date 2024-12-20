// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package ecode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrSuccess-0]
	_ = x[ErrBadRequest-400]
	_ = x[ErrUnauthorized-401]
	_ = x[ErrNotFound-404]
	_ = x[ErrInternalServer-500]
	_ = x[ErrEmailAlreadyUse-10006]
	_ = x[ErrPasswordIncorrect-10007]
	_ = x[ErrUserNotFound-10008]
	_ = x[ErrAdminUserCanNotModify-10009]
	_ = x[ErrRoleHasUser-10010]
	_ = x[ErrRoleNotFound-10011]
	_ = x[ErrMenuHasChild-10012]
	_ = x[ErrMenuParentedNotFound-10013]
	_ = x[ErrDeptNotFound-10014]
	_ = x[ErrDeptHasChild-10015]
	_ = x[ErrDeptHasUser-10016]
	_ = x[ErrDeptParentNotFound-10017]
	_ = x[ErrAcmeIdNotFound-10018]
	_ = x[ErrAcmePathNotFound-10019]
	_ = x[ErrAcmeEmailNotFound-10020]
}

const (
	_ErrCode_name_0 = "Success"
	_ErrCode_name_1 = "Bad RequestUnauthorized"
	_ErrCode_name_2 = "Not Found"
	_ErrCode_name_3 = "Internal Server Error"
	_ErrCode_name_4 = "The email is already in useThe password is incorrectThe user does not existThe super administrator role cannot be modifiedThe role has users and cannot be deletedThe role not foundThe menu has children and cannot be deletedThe parent menu not foundThe department not foundThe department has children and cannot be deletedThe department has user and cannot be deletedThe parent department not foundThe acme id not foundThe acme path not foundThe acme email not found"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 11, 23}
	_ErrCode_index_4 = [...]uint16{0, 27, 52, 75, 122, 162, 180, 223, 248, 272, 321, 366, 397, 418, 441, 465}
)

func (i ErrCode) String() string {
	switch {
	case i == 0:
		return _ErrCode_name_0
	case 400 <= i && i <= 401:
		i -= 400
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case i == 404:
		return _ErrCode_name_2
	case i == 500:
		return _ErrCode_name_3
	case 10006 <= i && i <= 10020:
		i -= 10006
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
