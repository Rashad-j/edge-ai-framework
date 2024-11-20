package config

import "github.com/caarlos0/env"

type Config struct {
	ONNXModelPath       string `env:"ONNX_MODEL_PATH" envDefault:"../../model_examples/onnx.onnx"`
	TensorFlowModelPath string `env:"TENSOR_FLOW_MODEL_PATH" envDefault:"../../model_examples/tensorflow.tensorflow"`
	Port                string `env:"PORT" envDefault:":8083"`
}

func ReadEnvConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
