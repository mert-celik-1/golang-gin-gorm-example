package exceptions

type BadRequestException struct {
	Exception CustomException
	Error     error
}

func (e BadRequestException) CreateNew(err error) ICustomException {

	ce := CustomException{
		Success: false,
		Code:    BadRequestException_Code,
		Message: BadRequestException_Message,
	}

	return &AuthException{
		Exception: ce,
		Error:     err,
	}
}

func (e BadRequestException) GetCustomException() CustomException {
	return e.Exception
}
func (e BadRequestException) GetError() error {
	return e.Error
}
