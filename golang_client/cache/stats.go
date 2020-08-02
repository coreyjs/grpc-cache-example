package cache

import (
	"time"
)

// Stats - simple struct to hold start/end values
type Stats struct {
	StartedAt  time.Time
	FinishedAt time.Time
}
