package models

import (
	"time"
)

type Category struct {
	ID        int
	Name      string
	Slug      string
	CreatedAt time.Time
} 