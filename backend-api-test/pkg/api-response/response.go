package api_response

import "net/http"

func StatusOK(meta interface{}, message string) *Response {
	result := Response{
		Meta:    meta,
		Message: message,
		Status:  1,
		Code:    http.StatusOK,
		Data:    nil,
	}
	return &result
}

func Success(meta interface{}, data interface{}) *Response {
	result := Response{
		Meta:    meta,
		Message: "Success",
		Status:  1,
		Code:    http.StatusOK,
		Data:    data,
	}
	return &result
}

func BadRequest(errorMessages interface{}) *Response {
	result := Response{
		Meta:    nil,
		Message: errorMessages,
		Status:  0,
		Code:    http.StatusBadRequest,
		Data:    nil,
	}
	return &result
}

func UnAuthorized(meta interface{}, errorMessage string) *Response {
	result := Response{
		Meta:    meta,
		Message: errorMessage,
		Status:  0,
		Code:    http.StatusUnauthorized,
		Data:    nil,
	}
	return &result

}

func Forbidden(meta interface{}, message string) *Response {
	result := Response{
		Meta:    meta,
		Message: message + " Forbidden",
		Status:  0,
		Code:    http.StatusForbidden,
		Data:    nil,
	}
	return &result
}

func DataNotFound() *Response {
	result := Response{
		Meta:    nil,
		Message: "Data Not Found",
		Status:  0,
		Code:    http.StatusNotFound,
		Data:    nil,
	}
	return &result
}

func NotFound(meta interface{}, message string) *Response {
	result := Response{
		Meta:    meta,
		Message: message,
		Status:  0,
		Code:    http.StatusNotFound,
		Data:    nil,
	}
	return &result
}

func NotAllowed(meta interface{}, message string) *Response {
	result := Response{
		Meta:    meta,
		Message: message + " Not Allowed",
		Status:  0,
		Code:    http.StatusMethodNotAllowed,
		Data:    nil,
	}
	return &result
}

func ServerError(errorMessages interface{}) *Response {
	result := Response{
		Meta:    nil,
		Message: errorMessages,
		Status:  0,
		Code:    http.StatusInternalServerError,
		Data:    nil,
	}
	return &result
}

func ServerNotImplemented(errorMessages interface{}) *Response {
	result := Response{
		Meta:    nil,
		Message: errorMessages,
		Status:  0,
		Code:    http.StatusNotImplemented,
		Data:    nil,
	}
	return &result
}
