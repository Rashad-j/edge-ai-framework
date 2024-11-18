package ai

import (
	"fmt"

	"github.com/Rashad-j/edge-ai-framework/pkg/ai/runtime"
)

const (
	ONNX   = "onnx"
	TFLite = "tflite"
)

type AIModelFacade struct {
	onnxRuntime   *runtime.ONNXRuntime
	tfliteRuntime *runtime.TFLiteRuntime
	loadedRuntime runtime.Runtime
}

// NewAIModelFacade creates a new instance of the AIModelFacade
func NewAIModelFacade(ONNXPath, TFPath string) *AIModelFacade {
	return &AIModelFacade{
		onnxRuntime:   runtime.NewONNXRunTime(ONNXPath),
		tfliteRuntime: runtime.NewTFLiteRuntTime(TFPath),
	}
}

// LoadModel loads the appropriate model based on file format
func (f *AIModelFacade) LoadModel(path string) error {
	switch path {
	case ONNX:
		f.loadedRuntime = f.onnxRuntime
	case TFLite:
		f.loadedRuntime = f.tfliteRuntime
	default:
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
