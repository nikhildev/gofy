package models

import "time"

type Change struct {
	ID            string
	ChangeTime    time.Time
	ChangeField   string
	ChangedBy     string
	PreviousValue string
	NewValue      string
}
