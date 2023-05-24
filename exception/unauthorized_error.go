package exception

type UnauthorizedError struct{
	Error string
}

func NewUserUnauthorized(err string) UnauthorizedError {
	return UnauthorizedError{Error: err}
}