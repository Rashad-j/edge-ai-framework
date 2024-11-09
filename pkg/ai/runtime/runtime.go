package runtime

type Runtime interface {
    LoadModel(path string) error
    Infer(input []float32) ([]float32, error)
}