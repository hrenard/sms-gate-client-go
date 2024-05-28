package smsgateway

import "errors"

// Processing state
type ProcessingState string

var ErrConflictFields = errors.New("conflict fields")
