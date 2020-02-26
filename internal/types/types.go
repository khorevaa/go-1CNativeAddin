package types

const (
	S_OK int = 0

	S_FALSE      int     = 1
	E_POINTER    uintptr = 0x80004003
	E_FAIL       uintptr = 0x80004005
	E_UNEXPECTED uintptr = 0x8000FFFF

	ADDIN_E_NONE             int = 1000
	ADDIN_E_ORDINARY             = 1001
	ADDIN_E_ATTENTION            = 1002
	ADDIN_E_IMPORTANT            = 1003
	ADDIN_E_VERY_IMPORTANT       = 1004
	ADDIN_E_INFO                 = 1005
	ADDIN_E_FAIL                 = 1006
	ADDIN_E_MSGBOX_ATTENTION     = 1007
	ADDIN_E_MSGBOX_INFO          = 1008
	ADDIN_E_MSGBOX_FAIL          = 1009
)
