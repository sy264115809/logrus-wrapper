package logrush

import (
	"github.com/Sirupsen/logrus"
)

type prefixHook struct {
	disableColor bool
}

var _ logrus.Hook = &prefixHook{}

func (h *prefixHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *prefixHook) Fire(entry *logrus.Entry) error {
	// records prefix as metadata (prefixed.TextFormatter will drop prefix anyway)
	if prefix, ok := entry.Data["prefix"]; ok && h.disableColor {
		entry.Data["@prefix"] = prefix
	}

	return nil
}
