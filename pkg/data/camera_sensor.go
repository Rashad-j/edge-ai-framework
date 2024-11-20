package data

import (
	"context"
	"fmt"
	"gocv.io/x/gocv"
	"time"
)

// CameraSensor represents a camera device.
type CameraSensor struct {
	DeviceID int // Camera device ID (e.g., 0 for default camera)
}

// StreamData continuously captures frames from the camera and sends them to the channel.
func (c CameraSensor) StreamData(ctx context.Context, ch chan<- SensorData) {
	webcam, err := gocv.OpenVideoCapture(c.DeviceID)
	if err != nil {
		fmt.Printf("Error opening camera: %v\n", err)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Println("Starting Camera Sensor streaming...")

	ticker := time.NewTicker(100 * time.Millisecond) // 10 FPS
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping Camera Sensor streaming...")
			return
		case <-ticker.C:
			if ok := webcam.Read(&img); !ok {
				fmt.Println("Error reading frame from camera.")
				continue
			}

			if img.Empty() {
				fmt.Println("Empty frame received.")
				continue
			}

			// Send the frame to the channel
			ch <- SensorData{Image: img.Clone(), Timestamp: time.Now()}
		}
	}
}
