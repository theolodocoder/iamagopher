package main

import (
	"log/syslog"
)

func sysLog() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
}
