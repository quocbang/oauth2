package errors

import "fmt"

type Error struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  Code   `json:"error_code"`
	Details    string `json:"details"`
	Raw        error  `json:"raw"`
}

func (e Error) Error() string {
	msg := e.Details
	if e.ErrorCode != 0 {
		msg = fmt.Sprintf("code: %d, details: %s", e.ErrorCode, msg)
	}
	if e.Raw != nil {
		msg = fmt.Sprintf("%s, raw: %v", msg, e.Raw)
	}
	return msg
}
