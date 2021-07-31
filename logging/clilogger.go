package logging

import "fmt"

// cliLogger is an implementation of the Logger
// which logs into the cli
type cliLogger struct {
}

func NewCLILogger() Logger {
	return &cliLogger{}
}

func (c *cliLogger) Log(log string)  {
	fmt.Println(log)
}

func (c *cliLogger) Logf(formatString string, params ...interface{}) {
	fmt.Printf(formatString, params...)
}
