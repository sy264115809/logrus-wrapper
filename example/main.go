package main

import "github.com/sy264115809/logrush"

func main() {
	defaultLogger()
	outputToFileLogger()
	verboseLogger()
	showCallerLogger()
	showCallerWithAdjustLogger()
	withoutColorLogger()
}

func defaultLogger() {
	logrush.New(&logrush.Config{}).WithFields(logrush.Fields{
		"output":      "stderr",
		"level":       "info",
		"show_prefix": "false",
		"color":       "true",
		"show_caller": "false",
	}).Info("default logger")
}

func outputToFileLogger() {
	logrush.New(&logrush.Config{
		Prefix: "outputToFileLogger",
		Output: "example.log",
	}).WithFields(logrush.Fields{
		"output":      "example.log in current directory",
		"level":       "info",
		"show_prefix": "true",
		"color":       "true",
		"show_caller": "false",
	}).Warnln("output to file logger")
}

func verboseLogger() {
	logrush.New(&logrush.Config{
		Prefix:  "verboseLogger",
		Output:  "example.log",
		Verbose: true,
		Level:   "debug",
	}).WithFields(logrush.Fields{
		"output":      "both stderr and example.log in current directory",
		"level":       "debug",
		"show_prefix": "true",
		"color":       "true",
		"show_caller": "false",
	}).Debugf("verbose logger")
}

func showCallerLogger() {
	logrush.New(&logrush.Config{
		Prefix:     "showCallerLogger",
		ShowCaller: true,
	}).WithFields(logrush.Fields{
		"output":      "stderr",
		"level":       "info",
		"show_prefix": "true",
		"color":       "true",
		"show_caller": "true",
	}).Error("show caller logger")
}

func showCallerWithAdjustLogger() {
	logrush.New(&logrush.Config{
		Prefix:            "showCallerWithAdjustLogger",
		ShowCaller:        true,
		CallerDepthAdjust: 1,
	}).WithFields(logrush.Fields{
		"output":      "stderr",
		"level":       "info",
		"show_prefix": "true",
		"color":       "true",
		"show_caller": "true",
	}).Info("show caller with adjust logger, will show the caller is main.go:<line in main()>")
}

func withoutColorLogger() {
	logrush.New(&logrush.Config{
		Prefix:        "withoutColorLogger",
		DisableColors: true,
		ShowCaller:    true,
	}).WithFields(logrush.Fields{
		"output":      "stderr",
		"level":       "info",
		"show_prefix": "true",
		"color":       "false",
		"show_caller": "true",
	}).Info("without color logger")
}
