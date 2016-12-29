package logrush

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// Config type.
type Config struct {
	Output       string
	Level        string
	Verbose      bool
	outputWirter io.Writer

	DisableColors bool
	Prefix        string

	ShowCaller        bool
	CallerDepthAdjust uint
}

// OutputWriter returns a io.Writer according to the config.
func (c *Config) OutputWriter() io.Writer {
	if c.outputWirter == nil {
		switch output := strings.ToLower(c.Output); output {
		case "":
			c.outputWirter = os.Stderr
		default:
			logPath := path.Clean(output)
			logDir := path.Dir(output)
			err := os.MkdirAll(logDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Get info of log directory(%s) fail, use os.Stderr instead: %s\n", logDir, err)
				c.outputWirter = os.Stderr
			} else {
				f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					fmt.Printf("Open log file(%s) fail, use os.Stderr instead: %s\n", logPath, err)
					c.outputWirter = os.Stderr
				} else {
					c.outputWirter = f
				}
			}
		}

		if c.Verbose && c.outputWirter != os.Stderr {
			c.outputWirter = io.MultiWriter(c.outputWirter, os.Stderr)
		}
	}
	return c.outputWirter
}
