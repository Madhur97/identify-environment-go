package detect

// Detector defines the interface for detecting an environment
type Detector interface {
	Detect() bool // Detect determines if it matches the current environment
	Info() string // Info returns a description of the environment
}
