// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Priority string

const (
	PriorityNone   Priority = "none"
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (e *Priority) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Priority(s)
	case string:
		*e = Priority(s)
	default:
		return fmt.Errorf("unsupported scan type for Priority: %T", src)
	}
	return nil
}

type NullPriority struct {
	Priority Priority
	Valid    bool // Valid is true if Priority is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPriority) Scan(value interface{}) error {
	if value == nil {
		ns.Priority, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Priority.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPriority) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Priority), nil
}

type Tasks struct {
	ID          uuid.UUID
	Description string
	Priority    Priority
	StartDate   pgtype.Timestamp
	DueDate     pgtype.Timestamp
	Done        bool
}
