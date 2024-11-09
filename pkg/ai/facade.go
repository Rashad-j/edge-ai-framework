package ai

import (
    "fmt"
    "github.com/Rashad-j/edge-ai-framework/pkg/ai/runtime"
)

type AIModelFacade struct {
    onnxRuntime    *runtime.ONNXRuntime
    tfliteRuntime  *runtime.TFLiteRuntime
    loadedRuntime  runtime.Runtime
}

// NewAIModelFacade creates a new instance of the AIModelFacade
func NewAIModelFacade() *AIModelFacade {
    return &AIModelFacade{
        onnxRuntime:   &runtime.ONNXRuntime{},
        tfliteRuntime: &runtime.TFLiteRuntime{},
    }
}

// LoadModel loads the appropriate model based on file format
func (f *AIModelFacade) LoadModel(path string) error {
    if path == "onnx" {
        f.loadedRuntime = f.onnxRuntime
    } else if path == "tflite" {
        f.loadedRuntime = f.tfliteRuntime
    } else {
        return fmt.Errorf("unsupported model format")
    }
    return f.loadedRuntime.LoadModel(path)
}

// RunInference runs inference on the loaded model
func (f *AIModelFacade) RunInference(input []float32) ([]float32, error) {
    if f.loadedRuntime == nil {
        return nil, fmt.Errorf("no model loaded")
    }
    return f.loadedRuntime.Infer(input)
}