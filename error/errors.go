package error

import (
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1"k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
	"net/http"
)

// StatusError defines a http error
type StatusError struct {
	// Status of the operation.
	// One of: "Success" or "Failure".
	Status string `json:"status,omitempty"`
	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty"`
	// A machine-readable description of why this operation is in the
	// "Failure" status. If this value is empty there
	// is no information available. A Reason clarifies an HTTP status
	// code but does not override it.
	Reason StatusReason `json:"reason,omitempty"`
	// Extended data associated with the reason.  Each reason may define its
	// own extended details. This field is optional and the data returned
	// is not guaranteed to conform to any schema except that defined by
	// the reason type.
	Details *StatusDetails `json:"details,omitempty"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code int32 `json:"code,omitempty"`
	// responsed data if success
	Data interface{} `json:"data,omitempty"`
}

// StatusDetails is a set of additional properties that MAY be set by the
// server to provide additional information about a response. The Reason
// field of a Status object defines what attributes will be set. Clients
// must ignore fields that do not match the defined type of each attribute,
// and should assume that any attribute may be empty, invalid, or under
// defined.
type StatusDetails struct {
	// The name attribute of the resource associated with the status StatusReason
	// (when there is a single name which can be described).
	Name string `json:"name,omitempty"`
	// The group attribute of the resource associated with the status StatusReason.
	Group string `json:"group,omitempty"`
	// The kind attribute of the resource associated with the status StatusReason.
	// On some operations may differ from the requested resource Kind.
	Kind string `json:"kind,omitempty"`
	// The Causes array includes more details associated with the StatusReason
	// failure. Not all StatusReasons may provide detailed causes.
	Causes []StatusCause `json:"causes,omitempty"`
	// If specified, the time in seconds before the operation should be retried.
	RetryAfterSeconds int32 `json:"retryAfterSeconds,omitempty"`
	// Edition level
	Level string `json:"level,omitempty"`
}

// Values of Status.Status
const (
	StatusSuccess = "Success"
	StatusFailure = "Failure"
)

// StatusReason is an enumeration of possible failure causes.  Each StatusReason
// must map to a single HTTP status code, but multiple reasons may map
// to the same HTTP status code.
// TODO: move to apiserver
type StatusReason string

const (
	// StatusReasonUnknown means the server has declined to indicate a specific reason.
	// The details field may contain other information about this error.
	// Status code 500.
	StatusReasonUnknown StatusReason = ""

	// StatusReasonUnauthorized means the server can be reached and understood the request, but requires
	// the user to present appropriate authorization credentials (identified by the WWW-Authenticate header)
	// in order for the action to be completed. If the user has specified credentials on the request, the
	// server considers them insufficient.
	// Status code 401
	StatusReasonUnauthorized StatusReason = "Unauthorized"

	// StatusReasonForbidden means the server can be reached and understood the request, but refuses
	// to take any further action.  It is the result of the server being configured to deny access for some reason
	// to the requested resource by the client.
	// Details (optional):
	//   "kind" string - the kind attribute of the forbidden resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the forbidden resource
	// Status code 403
	StatusReasonForbidden StatusReason = "Forbidden"



	// Status code 451 for StatusUnavailableForLegalReasons
	StatusLicenseExpire StatusReason = "LicenseExpire"
	// StatusPaymentRequired payment required
	StatusPaymentRequired StatusReason = "PaymentRequired"

	// StatusReasonNotFound means one or more resources required for this operation
	// could not be found.
	// Details (optional):
	//   "kind" string - the kind attribute of the missing resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the missing resource
	// Status code 404
	StatusReasonNotFound StatusReason = "NotFound"

	// StatusNotAcceptable means request params are not valid.
	// Details (optional):
	//   "kind" string - the kind attribute of the missing resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the missing resource
	// Status code 406
	StatusNotAcceptable StatusReason = "StatusNotAcceptable"

	// StatusReasonAlreadyExists means the resource you are creating already exists.
	// Details (optional):
	//   "kind" string - the kind attribute of the conflicting resource
	//   "id"   string - the identifier of the conflicting resource
	// Status code 409
	StatusReasonAlreadyExists StatusReason = "AlreadyExists"

	// StatusReasonConflict means the requested operation cannot be completed
	// due to a conflict in the operation. The client may need to alter the
	// request. Each resource may define custom details that indicate the
	// nature of the conflict.
	// Status code 409
	StatusReasonConflict StatusReason = "Conflict"

	// StatusReasonGone means the item is no longer available at the server and no
	// forwarding address is known.
	// Status code 410
	StatusReasonGone StatusReason = "Gone"

	// StatusPaymentRequired payment required
	StatusPreconditionFailed StatusReason = "PreconditionFailed"
	// StatusReasonInvalid means the requested create or update operation cannot be
	// completed due to invalid data provided as part of the request. The client may
	// need to alter the request. When set, the client may use the StatusDetails
	// message field as a summary of the issues encountered.
	// Details (optional):
	//   "kind" string - the kind attribute of the invalid resource
	//   "id"   string - the identifier of the invalid resource
	//   "causes"      - one or more StatusCause entries indicating the data in the
	//                   provided resource that was invalid.  The code, message, and
	//                   field attributes will be set.
	// Status code 422
	StatusReasonInvalid StatusReason = "Invalid"

	// StatusReasonServerTimeout means the server can be reached and understood the request,
	// but cannot complete the action in a reasonable time. The client should retry the request.
	// This is may be due to temporary server load or a transient communication issue with
	// another server. Status code 500 is used because the HTTP spec provides no suitable
	// server-requested client retry and the 5xx class represents actionable errors.
	// Details (optional):
	//   "kind" string - the kind attribute of the resource being acted on.
	//   "id"   string - the operation that is being attempted.
	//   "retryAfterSeconds" int32 - the number of seconds before the operation should be retried
	// Status code 500
	StatusReasonServerTimeout StatusReason = "ServerTimeout"

	// StatusReasonTimeout means that the request could not be completed within the given time.
	// Clients can get this response only when they specified a timeout param in the request,
	// or if the server cannot complete the operation within a reasonable amount of time.
	// The request might succeed with an increased value of timeout param. The client *should*
	// wait at least the number of seconds specified by the retryAfterSeconds field.
	// Details (optional):
	//   "retryAfterSeconds" int32 - the number of seconds before the operation should be retried
	// Status code 504
	StatusReasonTimeout StatusReason = "Timeout"

	// StatusReasonBadRequest means that the request itself was invalid, because the request
	// doesn't make any sense, for example deleting a read-only object.  This is different than
	// StatusReasonInvalid above which indicates that the API call could possibly succeed, but the
	// data was invalid.  API calls that return BadRequest can never succeed.
	StatusReasonBadRequest StatusReason = "BadRequest"

	// StatusReasonMethodNotAllowed means that the action the client attempted to perform on the
	// resource was not supported by the code - for instance, attempting to delete a resource that
	// can only be created. API calls that return MethodNotAllowed can never succeed.
	StatusReasonMethodNotAllowed StatusReason = "MethodNotAllowed"

	// StatusReasonInternalError indicates that an internal error occurred, it is unexpected
	// and the outcome of the call is unknown.
	// Details (optional):
	//   "causes" - The original error
	// Status code 500
	StatusReasonInternalError StatusReason = "InternalError"

	// StatusReasonExpired indicates that the request is invalid because the content you are requesting
	// has expired and is no longer available. It is typically associated with watches that can't be
	// serviced.
	// Status code 410 (gone)
	StatusReasonExpired StatusReason = "Expired"

	// StatusReasonServiceUnavailable means that the request itself was valid,
	// but the requested service is unavailable at this time.
	// Retrying the request after some time might succeed.
	// Status code 503
	StatusReasonServiceUnavailable StatusReason = "ServiceUnavailable"
)

// StatusCause provides more information about an api.Status failure, including
// cases when multiple errors are encountered.
type StatusCause struct {
	// A machine-readable description of the cause of the error. If this value is
	// empty there is no information available.
	Type CauseType `json:"reason,omitempty" protobuf:"bytes,1,opt,name=reason,casttype=CauseType"`
	// A human-readable description of the cause of the error.  This field may be
	// presented as-is to a reader.
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	// The field of the resource that has caused this error, as named by its JSON
	// serialization. May include dot and postfix notation for nested attributes.
	// Arrays are zero-indexed.  Fields may appear more than once in an array of
	// causes due to fields having multiple errors.
	// Optional.
	//
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	Field string `json:"field,omitempty" protobuf:"bytes,3,opt,name=field"`
}

// CauseType is a machine readable value providing more detail about what
// occurred in a status response. An operation may have multiple causes for a
// status (whether Failure or Success).
type CauseType string

const (
	// CauseTypeFieldValueNotFound is used to report failure to find a requested value
	// (e.g. looking up an ID).
	CauseTypeFieldValueNotFound CauseType = "FieldValueNotFound"
	// CauseTypeFieldValueRequired is used to report required values that are not
	// provided (e.g. empty strings, null values, or empty arrays).
	CauseTypeFieldValueRequired CauseType = "FieldValueRequired"
	// CauseTypeFieldValueDuplicate is used to report collisions of values that must be
	// unique (e.g. unique IDs).
	CauseTypeFieldValueDuplicate CauseType = "FieldValueDuplicate"
	// CauseTypeFieldValueInvalid is used to report malformed values (e.g. failed regex
	// match).
	CauseTypeFieldValueInvalid CauseType = "FieldValueInvalid"
	// CauseTypeFieldValueNotSupported is used to report valid (as per formatting rules)
	// values that can not be handled (e.g. an enumerated string).
	CauseTypeFieldValueNotSupported CauseType = "FieldValueNotSupported"
	// CauseTypeUnexpectedServerResponse is used to report when the server responded to the client
	// without the expected return type. The presence of this cause indicates the error may be
	// due to an intervening proxy or the server software malfunctioning.
	CauseTypeUnexpectedServerResponse CauseType = "UnexpectedServerResponse"
)

func (err *StatusError) Error() string {
	return err.Message
}

// 尝试将k8s api返回的错误转成404错误，转换失败则返回nil, false。
func ConvToNotFoundErr(name, kind string, err error) (*StatusError, bool) {
	if k8sErr, ok := err.(*k8serrors.StatusError); true == ok && http.StatusNotFound == k8sErr.ErrStatus.Code {
		return NewNotFoundErr(name, kind), true
	}
	return nil, false
}

// 尝试将k8s api返回的错误转成403错误，转换失败则返回nil, false。
func ConvToConflictErr(name, kind string, err error) (*StatusError, bool) {
	if k8sErr, ok := err.(*k8serrors.StatusError); true == ok &&
		http.StatusConflict == k8sErr.ErrStatus.Code &&
		meta_v1.StatusReasonAlreadyExists == k8sErr.ErrStatus.Reason {
		return NewConflictErr(name, kind, fmt.Errorf("resource already exists")), true
	}
	return nil, false
}

// 尝试将k8s api返回的错误转成超时错误，转换失败则返回nil, false。
func ConvToTimeoutErr(kind, operation string, err error) (*StatusError, bool) {
	if k8sErr, ok := err.(*k8serrors.StatusError); true == ok &&
		http.StatusInternalServerError == k8sErr.ErrStatus.Code &&
		meta_v1.StatusReasonServerTimeout == k8sErr.ErrStatus.Reason {
		return NewTimeoutErr(kind, operation, 5), true
	}
	return nil, false
}

func NewSuccessStatus(results interface{}) *StatusError {
	return &StatusError{
		Data:   results,
		Status: StatusSuccess,
		Code:   http.StatusOK,
	}
}

func NewPaymentErr(kind, message string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusPaymentRequired,
		Reason: StatusPaymentRequired,
		Details: &StatusDetails{
			Kind: kind,
		},
		Message: message,
	}
}

func NewForbiddenErr(kind, name, message string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusForbidden,
		Reason: StatusReasonForbidden,
		Details: &StatusDetails{
			Kind: kind,
			Name: name,
		},
		Message: message,
	}
}
func NewLicenseExpireErr(kind string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusUnavailableForLegalReasons,
		Reason: StatusLicenseExpire,
		Details: &StatusDetails{
			Kind: kind,
		},
		Message: "License Expired",
	}
}

func NewResourceIsNotEnough(kind, message string, data interface{}) *StatusError {
	return &StatusError {
		Status: StatusFailure,
		Code: http.StatusForbidden,
		Reason: StatusReasonForbidden,
		Data: data,
		Details: &StatusDetails{
			Kind: kind,
		},
		Message: "The cluster resource is not enough",
	}
}


func NewNotFoundErr(name, kind string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusNotFound,
		Reason: StatusReasonNotFound,
		Details: &StatusDetails{
			Kind: kind,
			Name: name,
		},
		Message: fmt.Sprintf("%s %s not found", kind, name),
	}
}

func NewBadRequestErr(msg string, details *StatusDetails) *StatusError {
	return &StatusError{
		Status:  StatusFailure,
		Code:    http.StatusBadRequest,
		Reason:  StatusReasonBadRequest,
		Message: msg,
		Details: details,
	}
}

func NewUnprocessableErr(msg string, details *StatusDetails) *StatusError {
	return &StatusError{
		Status:  StatusFailure,
		Code:    http.StatusUnprocessableEntity,
		Reason:  "UnprocessableEntity",
		Message: msg,
		Details: details,
	}
}

func NewUnauthorizedErr() *StatusError {
	message := "not authorized"
	return &StatusError{
		Status:  StatusFailure,
		Code:    http.StatusUnauthorized,
		Reason:  StatusReasonUnauthorized,
		Message: message,
	}
}

func NewNotAcceptableErr(name, kind string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusNotAcceptable,
		Reason: StatusNotAcceptable,
		Details: &StatusDetails{
			Kind: kind,
			Name: name,
		},
		Message: fmt.Sprintf("Invalid request for %s %s", kind, name),
	}
}

func NewAlreadyExistErr(name, kind string, err error) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusConflict,
		Reason: StatusReasonAlreadyExists,
		Details: &StatusDetails{
			Kind: kind,
			Name: name,
		},
		Message: fmt.Sprintf("Operation cannot be fulfilled on %s %q: %v", kind, name, err),
	}
}

func NewConflictErr(name, kind string, err error) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusConflict,
		Reason: StatusReasonConflict,
		Details: &StatusDetails{
			Kind: kind,
			Name: name,
		},
		Message: fmt.Sprintf("Operation cannot be fulfilled on %s %q: %v", kind, name, err),
	}
}

func NewPreconditionFailedErr(kind, level, message string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusPreconditionFailed,
		Reason: StatusPreconditionFailed,
		Details: &StatusDetails{
			Kind:  kind,
			Level: level,
		},
		Message: message,
	}
}

func NewInternalErr(err error) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusInternalServerError,
		Reason: StatusReasonInternalError,
		Details: &StatusDetails{
			Causes: []StatusCause{{Message: err.Error()}},
		},
		Message: fmt.Sprintf("Internal error occurred: %v", err),
	}
}

func NewTimeoutErr(kind, operation string, retryAfterSeconds int) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusInternalServerError,
		Reason: StatusReasonServerTimeout,
		Details: &StatusDetails{
			Kind:              kind,
			Name:              operation,
			RetryAfterSeconds: int32(retryAfterSeconds),
		},
		Message: fmt.Sprintf("The %s operation against %s could not be completed at this time, please try again.", operation, kind),
	}
}

func NewNoRegistryErr(result string) *StatusError {
	return &StatusError{
		Status: StatusFailure,
		Code:   http.StatusNoContent,
		Reason: StatusNotAcceptable,
		Message: result,
	}
}
