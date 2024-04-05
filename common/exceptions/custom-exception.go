package exceptions

type ICustomException interface {
	GetCustomException() CustomException
	GetError() error
}

type CustomException struct {
	Success bool
	Code    string
	Message string
}

const (
	Success_Code    = "100"
	Success_Message = "İşlem başarılı."

	AuthTokenProvided_Code    = "20"
	AuthTokenProvided_Message = "Auth hata meydana geldi"

	BadRequestException_Code    = "21"
	BadRequestException_Message = "Parametrelerde hata meydana geldi"

	NotFoundException_Code    = "22"
	NotFoundException_Message = "Veri bulunamadı"
)
