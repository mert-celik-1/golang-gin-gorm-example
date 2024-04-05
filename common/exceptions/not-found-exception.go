package exceptions

type NotFoundException struct {
	Exception CustomException
	Error     error
}

func (e NotFoundException) CreateNew(err error) ICustomException {

	ce := CustomException{
		Success: false,
		Code:    NotFoundException_Code,
		Message: NotFoundException_Message,
	}

	return &AuthException{
		Exception: ce,
		Error:     err,
	}
}

func (e NotFoundException) GetCustomException() CustomException {
	return e.Exception
}
func (e NotFoundException) GetError() error {
	return e.Error
}
