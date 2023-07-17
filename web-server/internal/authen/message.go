package authen

const (
	ErrBadRequest        = "Bad Request"
	ErrInternal          = "Internal Server Error"
	ErrNoSession         = "SessionID not found"
	ErrSessionExist      = "Current SessionID already exists"
	ErrNoPermission      = "This SessionID has no permission to do this"
	ErrDuplicateUsername = "The username has already been registered"
	ErrIncorrectUsername = "Incorrect Username or Password"
	ErrIncorrectPassword = "Incorrect Username or Password"
	ErrWrongEmail		 = "This email doesn't exists"
	MessageSuccessRegister = "Successful registration"
	MessageSuccessLogin    = "Successful login"
	MessageSuccessLogout   = "Successful logout"
	MessageChangePasswordSuccess = "Successful change password"
)
