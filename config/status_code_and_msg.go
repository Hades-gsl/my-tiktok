package config

const VideosCount = 30

const (
	OKStatusCode            = iota
	SQLQueryErrorStatusCode = 10000 + iota
	NoVideoStatusCode
	NameExistStatusCode
	NameOrPasswordEmptyStatusCode
	PasswordHashErrorStatusCode
	GenerateTokenErrorStatusCode
	SQLSaveErrorStatusCode
	UserNotFoundStatusCode
	PasswordWrongStatusCode
	NotLogInStatusCode
	IdInvalidStatusCode
	TokenInvalidStatusCode
	GetFileErrorStatusCode
	FileOpenErrorStatusCode
	FileReadErrorStatusCode
	FileTypeErrorStatusCode
	FileSaveErrorStatusCode
	TitleEmptyStatusCode
	NoIDStatusCode
	ParameterErrorStatusCode
	UnknownFavoriteTypeStatusCode
	SQLDeleteErrorStatusCode
)

var (
	OKStatusMsg                  = "OK"
	SQLQueryErrorStatusMsg       = "SQL query error"
	NoVideoStatusMsg             = "No video"
	NameExistStatusMsg           = "Name exists"
	NameOrPasswordEmptyStatusMsg = "Name or password is empty"
	PasswordHashErrorStatusMsg   = "Password hash error"
	GenerateTokenErrorStatusMsg  = "Generate token error"
	SQLSaveErrorStatusMsg        = "SQL save error"
	UserNotFoundStatusMsg        = "User not found"
	PasswordWrongStatusMsg       = "Password wrong"
	NotLogInStatusMsg            = "Not log in"
	IdInvalidStatusMsg           = "Id invalid status"
	TokenInvalidStatusMsg        = "Token invalid"
	GetFileErrorStatusMsg        = "Get file error"
	FileOpenErrorStatusMsg       = "File open error"
	FileReadErrorStatusMsg       = "File read error"
	FileTypeErrorStatusMsg       = "File type error"
	FileSaveErrorStatusMsg       = "File save error"
	TitleEmptyStatusMsg          = "Title empty"
	NoIDStatusMsg                = "No ID"
	ParameterErrorStatusMsg      = "Parameter error"
	UnknownFavoriteTypeStatusMsg = "Unknown favorite type"
	SQLDeleteErrorStatusMsg      = "SQL delete error"
)
