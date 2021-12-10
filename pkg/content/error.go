package content

import "errors"

var UserPwdError = errors.New("user input password error")

var SessionExpiredErrOR = errors.New("login info has been expired ,please reLogin")
