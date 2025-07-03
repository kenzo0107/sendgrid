package sendgrid

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	buf := bytes.NewBufferString("")
	logger := internalLog{logger: log.New(buf, "", 0|log.Lshortfile)}
	logger.Println("test line 123")
	assert.Equal(t, buf.String(), "logger_test.go:14: test line 123\n")
	buf.Truncate(0)
	logger.Print("test line 123")
	assert.Equal(t, buf.String(), "logger_test.go:17: test line 123\n")
	buf.Truncate(0)
	logger.Printf("test line 123\n")
	assert.Equal(t, buf.String(), "logger_test.go:20: test line 123\n")
	buf.Truncate(0)
	if err := logger.Output(1, "test line 123\n"); err != nil {
		log.Println(err)
	}
	assert.Equal(t, buf.String(), "logger_test.go:23: test line 123\n")
	buf.Truncate(0)
}

type errorLogger struct{}

func (e errorLogger) Output(calldepth int, s string) error {
	return assert.AnError
}

func TestLoggingWithError(t *testing.T) {
	// Override logFatal to capture calls instead of terminating the test
	originalLogFatal := logFatal
	defer func() { logFatal = originalLogFatal }()

	var fatalCalled bool
	logFatal = func(v ...interface{}) {
		fatalCalled = true
	}

	logger := internalLog{logger: errorLogger{}}

	// Test Println with error
	fatalCalled = false
	logger.Println("test")
	assert.True(t, fatalCalled, "logFatal should have been called for Println error")

	// Test Print with error
	fatalCalled = false
	logger.Print("test")
	assert.True(t, fatalCalled, "logFatal should have been called for Print error")

	// Test Printf with error
	fatalCalled = false
	logger.Printf("test")
	assert.True(t, fatalCalled, "logFatal should have been called for Printf error")
}
