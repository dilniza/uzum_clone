package models

// ResponseSuccess ...
type ResponseSuccess struct {
	Metadata interface{} `json:"metadata"`
	Data     interface{} `json:"data"`
}

// ResponseError ...
type ResponseError struct {
	Error interface{} `json:"error"`
}

// InternalServerError ...
type InternalServerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ValidationError ...
type ValidationError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	UserMessage string `json:"user_message"`
}

type ResponseOK struct {
	Message interface{} `json:"message"`
}

type Response struct {
	ID interface{} `json:"id"`
}

type ErrorWithDescription struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ErrorReason struct {
	Reason string `json:"reason"`
}

type ResponseResult struct {
	Result string `json:"result"`
}
