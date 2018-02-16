package response

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	// ErrNewResponse is error on newResponse
	ErrNewResponse = errors.New("app-interface-response: newResponse failed")
)

// Response is struct to represent HTTP Response
type Response struct {
	Body []byte
}

func newResponse(src interface{}) (*Response, error) {
	j, err := json.Marshal(src)

	if err != nil {
		return nil, errors.Wrap(err, ErrNewResponse.Error())
	}

	r := Response{Body: j}

	return &r, nil
}
