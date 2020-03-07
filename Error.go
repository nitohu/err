package err

type errorMsg struct {
	origin string
	msg    string
}

// Error is the basic error type
type Error struct {
	messages []errorMsg
}

// Init initializes the error type, adds the first msg
func (e *Error) Init(origin, msg string) {
	err := errorMsg{
		origin: origin,
		msg:    msg,
	}
	e.messages = []errorMsg{err}
}

// AddTraceback adds a traceback to the error message stack
func (e *Error) AddTraceback(origin, msg string) {
	origin = "Traceback:\n" + origin
	err := errorMsg{
		origin: origin,
		msg:    msg,
	}
	e.messages = append(e.messages, err)
}

// Error function which implements the error type
func (e Error) Error() string {
	var msg string

	for i := 0; i < len(e.messages); i++ {
		currentMsg := e.messages[i]
		errorMsg := currentMsg.origin + ": " + currentMsg.msg + "\n"

		msg += errorMsg
	}

	return msg
}

// Empty checks if the current Error object is empty, it returns true if the object is empty
func (e Error) Empty() bool {
	if e.messages == nil || len(e.messages) == 0 {
		return true
	}
	return false
}
