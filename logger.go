package logrush

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	DefaultTimestampFormat = ""
)

// Logger is the wrapper of logrus Logger
type Logger struct {
	*logrus.Logger

	prefix          string
	showCaller      bool
	callDepthOffset int
}

// Entry proxies the logrus.Entry to remove namespace dependency
type Entry struct {
	*logrus.Entry
}

// WithField is the proxy of entry.WithField
func (e *Entry) WithField(key string, value interface{}) *Entry {
	e.Entry = e.Entry.WithField(key, value)
	return e
}

// WithFields is the proxy of entry.WithFields
func (e *Entry) WithFields(fields Fields) *Entry {
	e.Entry = e.Entry.WithFields(logrus.Fields(fields))
	return e
}

// WithError is the proxy of entry.WithError
func (e *Entry) WithError(err error) *Entry {
	e.Entry = e.Entry.WithError(err)
	return e
}

// Fields holds the data should be output in log.
type Fields logrus.Fields

// New returns a logger with specified prefix.
func New(c *Config) *Logger {
	logger := &Logger{
		Logger: logrus.New(),
	}
	logger.Out = c.OutputWriter()

	if lvl, err := logrus.ParseLevel(c.Level); err == nil {
		logger.Level = lvl
	}

	logger.setupPrefix(c.Prefix, c.DisableColors)
	logger.setupCaller(c.ShowCaller, c.CallPathLength)

	return logger
}

//Prefix returns prefix
func (logger *Logger) Prefix() string {
	return logger.prefix
}

func (logger *Logger) setupPrefix(prefix string, disableColor bool) {
	logger.Formatter = &TextFormatter{
		DisableColors:   disableColor,
		ForceColors:     !disableColor,
		TimestampFormat: DefaultTimestampFormat,
	}
	logger.prefix = prefix
	logger.Hooks.Add(&prefixHook{
		disableColor: disableColor,
	})
}

func (logger *Logger) setupCaller(show bool, displayLength int) {
	logger.showCaller = show
	logger.Hooks.Add(&callerHook{
		show:          show,
		displayLength: displayLength,
	})
}

// Copy copies logger with new prefix if provided.
func (logger *Logger) Copy(prefix ...string) *Logger {
	p := logger.prefix
	if len(prefix) > 0 {
		p = prefix[0]
	}
	return &Logger{
		Logger:     logger.Logger,
		prefix:     p,
		showCaller: logger.showCaller,
	}
}

// SetCallDepthOffset sets call depth offset if config.ShowCaller is set.
func (logger Logger) SetCallDepthOffset(offset int) *Logger {
	logger.callDepthOffset = offset
	return &logger
}

// NewEntry returns an entry Entry
func (logger *Logger) entry() *Entry {
	entry := logrus.NewEntry(logger.Logger)
	if logger.prefix != "" {
		entry = entry.WithField("prefix", logger.prefix)
	}

	if logger.showCaller {
		entry = entry.WithField(keyCallDepthOffset, logger.callDepthOffset)
	}

	return &Entry{
		Entry: entry,
	}
}

// WithField adds a field to the log entry, note that you it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return logger.entry().WithField(key, value)
}

// WithFields adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	return logger.entry().WithFields(fields)
}

// WithError adds an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *Entry {
	return logger.entry().WithError(err)
}

// Debugf is the proxy of logger.Debugf()
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.entry().Debugf(format, args...)
}

// Infof is the proxy of logger.Infof()
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.entry().Infof(format, args...)
}

// Printf is the proxy of logger.Printf()
func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.entry().Printf(format, args...)
}

// Warnf is the proxy of logger.Warnf()
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.entry().Warnf(format, args...)
}

// Warningf is the proxy of logger.Warningf()
func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.entry().Warnf(format, args...)
}

// Errorf is the proxy of logger.Errorf()
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.entry().Errorf(format, args...)
}

// Fatalf is the proxy of logger.Fatalf()
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.entry().Fatalf(format, args...)
	os.Exit(1)
}

// Panicf is the proxy of logger.Panicf()
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.entry().Panicf(format, args...)
}

// Debug is the proxy of logger.Debug()
func (logger *Logger) Debug(args ...interface{}) {
	logger.entry().Debug(args...)
}

// Info is the proxy of logger.Info()
func (logger *Logger) Info(args ...interface{}) {
	logger.entry().Info(args...)
}

// Print is the proxy of logger.Print()
func (logger *Logger) Print(args ...interface{}) {
	logger.entry().Info(args...)
}

// Warn is the proxy of logger.Warn()
func (logger *Logger) Warn(args ...interface{}) {
	logger.entry().Warn(args...)
}

// Warning is the proxy of logger.Warning()
func (logger *Logger) Warning(args ...interface{}) {
	logger.entry().Warn(args...)
}

// Error is the proxy of logger.Error()
func (logger *Logger) Error(args ...interface{}) {
	logger.entry().Error(args...)
}

// Fatal is the proxy of logger.Fatal()
func (logger *Logger) Fatal(args ...interface{}) {
	logger.entry().Fatal(args...)
	os.Exit(1)
}

// Panic is the proxy of logger.Panic()
func (logger *Logger) Panic(args ...interface{}) {
	logger.entry().Panic(args...)
}

// Debugln is the proxy of logger.Debugln()
func (logger *Logger) Debugln(args ...interface{}) {
	logger.entry().Debugln(args...)
}

// Infoln is the proxy of logger.Infoln()
func (logger *Logger) Infoln(args ...interface{}) {
	logger.entry().Infoln(args...)
}

// Println is the proxy of logger.Println()
func (logger *Logger) Println(args ...interface{}) {
	logger.entry().Println(args...)
}

// Warnln is the proxy of logger.Warnln()
func (logger *Logger) Warnln(args ...interface{}) {
	logger.entry().Warnln(args...)
}

// Warningln is the proxy of logger.Warningln()
func (logger *Logger) Warningln(args ...interface{}) {
	logger.entry().Warnln(args...)
}

// Errorln is the proxy of logger.Errorln()
func (logger *Logger) Errorln(args ...interface{}) {
	logger.entry().Errorln(args...)
}

// Fatalln is the proxy of logger.Fatalln()
func (logger *Logger) Fatalln(args ...interface{}) {
	logger.entry().Fatalln(args...)
	os.Exit(1)
}

// Panicln is the proxy of logger.Panicln()
func (logger *Logger) Panicln(args ...interface{}) {
	logger.entry().Panicln(args...)
}
