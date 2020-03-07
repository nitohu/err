package err

type errorMsg struct {
	origin string
	msg    string
}

// Error is the basic error type
type Error struct {
	messages []errorMsg
}

func (e *accError) Init(origin, msg string) {
	err := errorMsg{
		origin: origin,
		msg:    msg,
	}
	e.messages = []errorMsg{err}
}

func (e *accError) AddTraceback(origin, msg string) {
	origin = "Traceback:\n" + origin
	err := errorMsg{
		origin: origin,
		msg:    msg,
	}
	e.messages = append(e.messages, err)
}

func (e accError) Error() string {
	var msg string

	for i := 0; i < len(e.messages); i++ {
		currentMsg := e.messages[i]
		errorMsg := currentMsg.origin + ": " + currentMsg.msg + "\n"

		msg += errorMsg
	}

	return msg
}
