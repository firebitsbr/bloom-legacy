package errors

type Code int

const (
	// Internal should be used for internal errors, where no other code fits
	Internal Code = iota
	// Not
	NotFound
	Unauthenticated
	PermissionDenied
)

type Error struct {
	Code    string
	Message string
}

func New(code Code, message string) Error {
	codeStr := "INTERNAL"

	switch code {
	case NotFound:
		codeStr = "NOT_FOUND"
	case PermissionDenied:
		codeStr = "PERMISSION_DENIED"
	case Unauthenticated:
		codeStr = "UNAUTHENTICATED"
	}
	return Error{
		Code:    codeStr,
		Message: message,
	}
}

func (err Error) Error() string {
	return err.Message
}
