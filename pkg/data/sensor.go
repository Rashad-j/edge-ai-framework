package data

import (
	"context"
	"gocv.io/x/gocv"
	"time"
)

type Sensor interface {
	StreamData(ctx context.Context, ch chan<- SensorData)
}

type SensorData struct {
	Image     gocv.Mat // Raw image data
	Timestamp time.Time
}
