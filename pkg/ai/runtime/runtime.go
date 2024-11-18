package runtime

// Runtine represents any runtime model, e.g. onnx or tensorflow
type Runtime interface {
    LoadModel(path string) error
    Infer(input []float32) ([]float32, error)
}