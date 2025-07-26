package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "C001"
	BadParameter        ErrCode = "C002"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidator         ErrCode = "A002"
	Unauthorizated              ErrCode = "A003"
)

// 元となるエラーを受け取ってMyAppError型にラップして返却する
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
