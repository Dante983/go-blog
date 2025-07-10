package models

import (
	"time"
)

type Tag struct {
	ID        int
	Name      string
	Slug      string
	CreatedAt time.Time
} 