package definition

type ResponseStyle struct {
	Code    string
	Message string
	Data    interface{}
}

const (
	/*Exception*/
	OK                   = "OK"
	BindError            = "BIND_ERROR"
	InvalidParam         = "INVALID_PARAM"
	UnknownError         = "UNKNOWN_ERROR"
	CreateRecordError    = "NEW_RECORD_ERROR"
	RetrieveRecordFailed = "RETRIEVE_RECORD_FAILED"
	UpdateRecordError    = "UPDATE_RECORD_ERROR"
	DeleteRecordError    = "NEW_RECORD_ERROR"
	RecordNoMatch        = "RECORD_NO_MATCH"
	SetCacheError        = "SET_CACHE_ERROR"
	GetCacheError        = "GET_CACHE_ERROR"
	ConnectionFailed     = "CONNECTION_FAILED"
	ResponseTimeOut      = "RESPONSE_TIMEOUT"
	RequestFailed        = "REQUEST_FAILED"
	UnmarshalFailed      = "UNMARSHAL_FAILED"
	MarshalFailed        = "MARSHAL_FAILED"
	PublishFailed        = "PUBLISH_FAILED"
	DecodeFailed         = "DECODE_FAILED"
	EmptyMsgBody         = "EMPTY_MSG_BODY"
	DeviceOffline        = "DEVICE_OFFLINE"
	SendRequestFailed    = "SEND_REQUEST_FAILED"
	DeviceResponseFailed = "DEVICE_RESPONSE_FAILED"
	CreateXLSXFailed     = "CREATE_XLSX_FAILED"
)
