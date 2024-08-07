package smsgateway

import "time"

type LogEntryPriority string

const (
	LogEntryPriorityDebug LogEntryPriority = "DEBUG"
	LogEntryPriorityInfo  LogEntryPriority = "INFO"
	LogEntryPriorityWarn  LogEntryPriority = "WARN"
	LogEntryPriorityError LogEntryPriority = "ERROR"
)

// LogEntry represents a log entry
type LogEntry struct {
	ID        uint64            `json:"id"`        // A unique identifier for the log entry.
	Priority  LogEntryPriority  `json:"priority"`  // The priority level of the log entry.
	Module    string            `json:"module"`    // The module or component of the system that generated the log entry.
	Message   string            `json:"message"`   // A message describing the log event.
	Context   map[string]string `json:"context"`   // Additional context information related to the log entry, typically including data relevant to the log event.
	CreatedAt time.Time         `json:"createdAt"` // The timestamp when this log entry was created.
}
