package barion

// ErrorResponse is the general error type
type ErrorResponse struct {
	status string
	Errors *[]barionError `json:"Errors"`
}

func (e *ErrorResponse) Error() string {
	if e.Errors != nil {
		val := e.status + ": "
		for _, s := range *e.Errors {
			val = val + s.ErrorCode
		}
		return val
	} else {
		return e.status
	}
}

type barionError struct {
	ErrorCode   string
	Title       string
	Description string
	//EndPoint    string
	AuthData   string
	HappenedAt string
}
