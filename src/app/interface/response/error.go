package response

// NewErrorResponse returns error response
func NewErrorResponse(message string) (*Response, error) {
	src := struct {
		Error string `json:"error"`
	}{message}

	return newResponse(src)
}
