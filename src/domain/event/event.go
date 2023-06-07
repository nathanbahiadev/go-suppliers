package event

import (
	"time"
)

type Event struct {
	StartedAt time.Time
	Data      interface{}
}
