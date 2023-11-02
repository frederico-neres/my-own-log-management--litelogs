package domain

import "time"

type LogDomain struct {
	Time    string
	Level   int
	Service string
	Message string
}

func NewLogDomain(level int, service, message string) *LogDomain {
	return &LogDomain{
		Time:    time.Now().Format(time.ANSIC),
		Level:   level,
		Service: service,
		Message: message,
	}
}
