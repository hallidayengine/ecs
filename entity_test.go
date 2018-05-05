package ecs_test

import (
	"testing"

	"github.com/hallidayengine/ecs"
	"github.com/segmentio/ksuid"
)

func TestIsActive(t *testing.T) {
	e := ecs.Entity{ID: ksuid.New(), Active: false}
	if e.IsActive() {
		t.Errorf("IsActive() = %t; e.Active = %t", e.IsActive(), false)
	}
}
