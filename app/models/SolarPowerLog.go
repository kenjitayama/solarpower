package models

import "time"

type SolarPowerLog struct {
	Id int64
	LogDate time.Time
	Generated float32
	Sold float32
	Bought float32
	Used float32
	P1Generated float32
	ExternalGenerated float32
	Income int32
	Outgo int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
