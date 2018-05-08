package model

import (
	"fmt"
)

// MarshalJSON renders trace id as a single hex string.
func (t TraceID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.AsString())), nil
}

// func (t *TraceID) UnmarshalJSON(data []byte) error {
// // TODO
// }

// MarshalJSON renders span id as a single hex string.
func (t SpanID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.AsString())), nil
}
