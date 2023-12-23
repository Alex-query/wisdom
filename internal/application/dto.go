package application

type RequestMessageGetSecurityCodeDTO struct {
	RequestMessageDTO
}

type RequestMessageGetWisdomDTO struct {
	RequestMessageDTO
}

type ResponseMessageGetWisdomDTO struct {
	ResponseMessageBaseDTO
	Data ResponseMessageGetWisdomDTOData `json:"data"`
}

type ResponseMessageGetWisdomDTOData struct {
	Wisdom string `json:"wisdom"`
}

type ResponseMessageGetSecurityCodeDTO struct {
	ResponseMessageBaseDTO
}

type RequestMessageDTO struct {
	Data    interface{}              `json:"data"`
	Command RequestMessageDTOCommand `json:"command"`
	Meta    RequestMetaDTO           `json:"meta"`
}

type RequestMessageDTOCommand string

const RequestMessageGetSecurityCodeDTOCommand RequestMessageDTOCommand = "get_security_code"
const RequestMessageGetWisdomDTOCommand RequestMessageDTOCommand = "get_wisdom"

type ErrorMessageDTO struct {
	ResponseMessageBaseDTO
	ErrorMessage string `json:"error_message"`
}

type ResponseMessageBaseDTO struct {
	Command RequestMessageDTOCommand `json:"command"`
	Meta    ResponseMetaDTO          `json:"meta"`
}

type ResponseMetaDTO struct {
	Code          int    `json:"code"`
	TaskToResolve string `json:"task_to_resolve"`
	RequestID     string `json:"request_id"`
}

type RequestMetaDTO struct {
	SecurityToken string `json:"security_token"`
	RequestID     string `json:"request_id"`
}
