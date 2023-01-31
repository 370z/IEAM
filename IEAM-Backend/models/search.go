package models

import "time"

type Search struct {
	Year  *time.Time `json:"year"`
	Month *time.Time `json:"month"`
	Type  *uint      `json:"type"`
}
