package exceptions

type AuthException struct {
	Exception CustomException
	Error     error
}

func (e AuthException) CreateNew(err error) ICustomException {

	ce := CustomException{
		Success: false,
		Code:    AuthTokenProvided_Code,
		Message: AuthTokenProvided_Message,
	}

	return &AuthException{
		Exception: ce,
		Error:     err,
	}
}

func (e AuthException) GetCustomException() CustomException {
	return e.Exception
}
func (e AuthException) GetError() error {
	return e.Error
}
