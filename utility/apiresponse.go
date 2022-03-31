package utility


// Response ...
type Response struct {
	Message string
	Status bool
}

// ResponseObj ...
type ResponseObj struct {
	Response
	Data interface{}
}

// ValidationResponseObj ...
type ValidationResponseObj struct {
	Response
	ValidationMsg string
}

// Success returns success response with data
func (r Response) Success(message string, data interface{}) ResponseObj {
	res:=Response{}
	res.Status= true
	res.Message= message

	return ResponseObj{
		Response: res,
		Data: data,
	}
}

//PlainSuccess returns success response without data
func (r Response) PlainSuccess(message string) Response {
	return Response{
		Status: true,
		Message: message,
	}
}

// Error returns error response
func (r Response) Error(message string) Response {
	return Response{
		Status: false,
		Message: message,
	}
}

// ValidationError return validation error
func (r Response) ValidationError(message, error string) ValidationResponseObj {
	res:=Response{
		Status: false,
		Message: message,
	}
	return ValidationResponseObj{
		Response:   res,
		ValidationMsg: error,
	}
}


// NewResponse ...
func NewResponse() Response {
	return Response{}
}

