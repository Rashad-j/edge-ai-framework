# Edge AI Framework - Golang Implementation

### Overview
The **Edge AI Framework** is an open-source project written in Golang designed to simplify the deployment of AI models on edge devices. It uses a **design pattern** to abstract away the complexities of working with multiple model formats (e.g., ONNX, TensorFlow Lite) and provides a simple, user-friendly interface for loading models and running inference. This makes it easy for developers to integrate AI capabilities into their IoT and edge applications.

## Features
- **Simplified Model Interaction**: Utilize the Facade Pattern to interact with AI models seamlessly without worrying about complex underlying logic.
- **Multiple Runtime Support**: Easily integrate ONNX and TensorFlow Lite models.
- **Modular Design**: Built with scalability in mind, adding new runtimes or extending functionalities is straightforward.
- **Logging and Config Management**: Centralized logging and configuration utilities for ease of use.

## Project Structure
Below is an overview of the project structure:

```
edge-ai-facade/
├── cmd/
│   ├── server/
│   │   └── main.go                 // Entry point for starting the Edge AI service
├── pkg/
│   ├── ai/
│   │   ├── facade.go               // Facade implementation providing a simple API for model interaction
│   │   ├── runtime/
│   │   │   ├── tflite_runtime.go   // TensorFlow Lite runtime integration
│   │   │   ├── onnx_runtime.go     // ONNX runtime integration
│   │   │   └── runtime.go          // Common runtime interface
│   ├── data/
│   │   ├── input.go                // Functions to handle input data from sensors or other sources
│   │   ├── preprocess.go           // Data preprocessing utilities
│   │   ├── postprocess.go          // Data postprocessing utilities
│   ├── config/
│   │   └── config.go               // Configuration management (e.g., model paths, environment variables)
│   ├── logger/
│   │   └── logger.go               // Centralized logging for the framework
├── api/
│   ├── server.go                   // REST/gRPC server to expose the AI functionalities via HTTP or gRPC
├── docs/
│   ├── README.md                   // Project overview and setup instructions
│   ├── architecture.md             // Documentation of system architecture and patterns used
├── test/
│   ├── facade_test.go              // Unit tests for facade layer
│   ├── inference_test.go           // Unit tests for inference runtimes
│   ├── integration_test.go         // Integration tests for end-to-end data handling
├── Dockerfile                      // Dockerfile for containerizing the application
├── go.mod                          // Go modules file for dependency management
├── go.sum                          // Checksums for module dependencies
└── Makefile                        // Common tasks like building, testing, and running the application
```

### Getting Started

#### Prerequisites
- **Golang**: Version 1.18 or higher is required to build the project.
- **ONNX Runtime**: The ONNX runtime must be installed for ONNX model inference.
- **TensorFlow Lite**: Install TensorFlow Lite libraries for TFLite model support.

#### Installation

1. **Clone the Repository**:
    ```sh
    git clone github.com/Rashad-j/edge-ai-framework.git
    cd edge-ai-facade
    ```

2. **Install Dependencies**:

3. **Run Example**:
    ```sh
    go run examples/basic_example.go
    ```

### Usage
The main component of this framework is the **AIModel** which provides simplified methods for interacting with AI models.

### Contributing
Contributions are welcome! Please create a pull request or open an issue for discussion.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Contact
Rashad Jamara


