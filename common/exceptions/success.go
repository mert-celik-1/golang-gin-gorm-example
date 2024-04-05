package exceptions

type Success struct {
	Exception CustomException
	Error     error
}

func (e Success) CreateNew(err error) ICustomException {

	ce := CustomException{
		Success: true,
		Code:    Success_Code,
		Message: Success_Message,
	}

	return &Success{
		Exception: ce,
		Error:     err,
	}
}

func (e Success) GetCustomException() CustomException {
	return e.Exception
}
func (e Success) GetError() error {
	return e.Error
}
