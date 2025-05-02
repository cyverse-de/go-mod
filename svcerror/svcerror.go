// Deprecated: Provided for compatibility with code that used to use protocol buffers.

package svcerror

import (
	"fmt"
	"maps"
	"slices"
)

// *
// The types of errors that can be retuned by message handlers.
type ErrorCode int32

const (
	ErrorCode_UNSET             ErrorCode = 0  // Default value for the error code. Don't set the error code to this. Use Unspecified if tempted.
	ErrorCode_UNSPECIFIED       ErrorCode = 1  // An error occurred, but the kind wasn't specified or included in the list.
	ErrorCode_INTERNAL          ErrorCode = 2  // Internal error.
	ErrorCode_NOT_FOUND         ErrorCode = 3  // The requested resource wasn't found.
	ErrorCode_BAD_REQUEST       ErrorCode = 4  // The request was bad/wrong is some way.
	ErrorCode_MARSHAL_FAILURE   ErrorCode = 5  // A failure to marshal a response.
	ErrorCode_UNMARSHAL_FAILURE ErrorCode = 6  // A failure to unmarshal a request.
	ErrorCode_PARAMETER_MISSING ErrorCode = 7  // A parameter is missing.
	ErrorCode_PARAMETER_INVALID ErrorCode = 8  /// A parameter is invalid.
	ErrorCode_UNAUTHENTICATED   ErrorCode = 9  /// Operation requires authentication, which was not provided.
	ErrorCode_FORBIDDEN         ErrorCode = 10 /// Operation is no allowed.
	ErrorCode_TIMEOUT           ErrorCode = 11 /// Operation timed out.
	ErrorCode_UNSUPPORTED       ErrorCode = 12 /// Operation is not supported.
	ErrorCode_UNIMPLEMENTED     ErrorCode = 13 /// Operation has not been implemented.
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "UNSET",
		1:  "UNSPECIFIED",
		2:  "INTERNAL",
		3:  "NOT_FOUND",
		4:  "BAD_REQUEST",
		5:  "MARSHAL_FAILURE",
		6:  "UNMARSHAL_FAILURE",
		7:  "PARAMETER_MISSING",
		8:  "PARAMETER_INVALID",
		9:  "UNAUTHENTICATED",
		10: "FORBIDDEN",
		11: "TIMEOUT",
		12: "UNSUPPORTED",
		13: "UNIMPLEMENTED",
	}
	ErrorCode_value = map[string]int32{
		"UNSET":             0,
		"UNSPECIFIED":       1,
		"INTERNAL":          2,
		"NOT_FOUND":         3,
		"BAD_REQUEST":       4,
		"MARSHAL_FAILURE":   5,
		"UNMARSHAL_FAILURE": 6,
		"PARAMETER_MISSING": 7,
		"PARAMETER_INVALID": 8,
		"UNAUTHENTICATED":   9,
		"FORBIDDEN":         10,
		"TIMEOUT":           11,
		"UNSUPPORTED":       12,
		"UNIMPLEMENTED":     13,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	v := int32(x)
	if slices.Contains(slices.Collect(maps.Keys(ErrorCode_name)), v) {
		return ErrorCode_name[v]
	}
	return fmt.Sprintf("ErrorCode(%d)", v)
}

func (x ErrorCode) Value() int32 {
	return int32(x)
}

// *
// An error returned by a request handler.
type ServiceError struct {
	// The numeric error code from the error code enum.
	ErrorCode ErrorCode `json:"error_code,omitempty"`
	// The status code for the error.
	StatusCode int32 `json:"status_code,omitempty"`
	// The error's message.
	Message string `json:"message,omitempty"`
}

func (x *ServiceError) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_UNSET
}

func (x *ServiceError) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ServiceError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}
