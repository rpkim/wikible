package api

// ErrorData is a format in error response from wiki api
type ErrorData struct {
	Authorized  bool		`json:"authorized"`
	Valid 		bool		`json:"valid"`
	Errors[]    string		`json:"errors"`
	Successful  bool		`json:"successful"`
}

// ErrorResponse is a struct for error response from wiki api
type ErrorResponse struct {
	StatusCode int			`json:"statusCode"`
	ErrorData				`json:"data"`
	Message string			`json:"message"`
}
