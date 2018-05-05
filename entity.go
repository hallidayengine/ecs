package ecs

import (
	"github.com/segmentio/ksuid"
)

// Entity is the base game object for everything in Halliday
type Entity struct {
	ID     ksuid.KSUID
	Active bool
}

// IsActive returns true or false for an active entity
// Since active and inactive entities will be divided
// into separate trees, you should not need to call this
// very often.
func (e Entity) IsActive() bool {
	return e.Active
}
