package backendTraineeAssignment2023

import "time"

type LogEntry struct {
	UserID    int       `db:"user_id"`
	SegmentID int       `db:"segment_id"`
	Action    string    `db:"action"`
	Datetime  time.Time `db:"datetime"`
}
