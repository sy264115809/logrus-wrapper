package logrush

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	keyCallAt                = "@at"
	keyCallDepthOffset       = "__call_depth_offset__"
	defaultDisplayPathLength = 3
)

type callerHook struct {
	show          bool
	displayLength int
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
	var depthOffset int
	if h.show {
		if o, ok := entry.Data[keyCallDepthOffset].(int); ok {
			depthOffset = o
		}

		for skip := 2; ; skip++ {
			if _, filePath, line, ok := runtime.Caller(skip); ok && !patternSkipCaller.MatchString(filePath) {
				if _, filePath, line, ok = runtime.Caller(skip + depthOffset); ok {
					parts := strings.SplitAfter(filePath, string(filepath.Separator))
					if l := h.DisplayLength(); l > 0 && len(parts) > l {
						filePath = strings.Join(parts[len(parts)-l:], "")
					}
					entry.Data[keyCallAt] = fmt.Sprintf("%s:%d", filePath, line)
					break
				}
			}
		}
		delete(entry.Data, keyCallDepthOffset)
	}
	return nil
}

func (h *callerHook) DisplayLength() int {
	if h.displayLength == 0 {
		return defaultDisplayPathLength
	}
	return h.displayLength
}
