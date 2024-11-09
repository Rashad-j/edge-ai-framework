package data

import "fmt"

type SensorData struct {
    Values []float32
}

// ReadSensorData simulates reading data from a sensor
func ReadSensorData() SensorData {
    fmt.Println("Reading data from sensor")
    // TODO: Replace with actual sensor reading logic
    return SensorData{Values: []float32{1.1, 2.2, 3.3}}
}