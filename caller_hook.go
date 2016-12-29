package logrush

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

type callerHook struct {
	depthAdjust int
}

var (
	patternSkipCaller *regexp.Regexp
)

func init() {
	patternSkipCaller, _ = regexp.Compile("logrush?/\\w+\\.go")
}

var _ logrus.Hook = &callerHook{}

func (h *callerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *callerHook) Fire(entry *logrus.Entry) error {
	for skip := 2; ; skip++ {
		if _, filePath, line, ok := runtime.Caller(skip); ok && !patternSkipCaller.MatchString(filePath) {
			if _, filePath, line, ok = runtime.Caller(skip + h.depthAdjust); ok {
				file := filepath.Base(filePath)
				parts := strings.SplitAfter(file, string(filepath.Separator))
				if len(parts) > 2 {
					file = strings.Join(parts[len(parts)-2:], "")
				}
				entry.Data["@at"] = fmt.Sprintf("%s:%d", file, line)
				break
			}
		}
	}
	return nil
}
