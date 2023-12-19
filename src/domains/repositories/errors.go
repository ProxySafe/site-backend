package repositories

type ErrDuplicateUser struct{}

func (e *ErrDuplicateUser) Error() string {
	return "user with such email/telephone/name already exists"
}

type ErrManyUsers struct{}

func (e *ErrManyUsers) Error() string {
	return "there are many users with such username"
}
