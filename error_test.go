package err

import (
	"fmt"
	"testing"
)

// TestError tests the Error() function of the err package
func TestError(t *testing.T) {
	var err Error
	err.Init("nonExistingFunc()", "Type Error")
	err.AddTraceback("TestError()", "Error while executing TestError()")
	msg := fmt.Sprint(err)
	correctMsg := "nonExistingFunc(): Type Error\nTraceback:\nTestError(): Error while executing TestError()\n"
	if msg != correctMsg {
		t.Errorf("Message was incorrect.\n-----\nError Message:\n%s\n-----\nCorrect Error Message:\n%s\n", msg, correctMsg)
	}
}
